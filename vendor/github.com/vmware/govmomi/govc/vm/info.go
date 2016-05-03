/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vm

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
	*flags.SearchFlag

	WaitForIP   bool
	General     bool
	ExtraConfig bool
	Resources   bool
}

func init() {
	cli.Register("vm.info", &info{})
}

func (cmd *info) Register(f *flag.FlagSet) {
	cmd.SearchFlag = flags.NewSearchFlag(flags.SearchVirtualMachines)

	f.BoolVar(&cmd.WaitForIP, "waitip", false, "Wait for VM to acquire IP address")
	f.BoolVar(&cmd.General, "g", true, "Show general summary")
	f.BoolVar(&cmd.ExtraConfig, "e", false, "Show ExtraConfig")
	f.BoolVar(&cmd.Resources, "r", false, "Show resource summary")
}

func (cmd *info) Process() error { return nil }

func (cmd *info) Run(f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	vms, err := cmd.VirtualMachines(f.Args())
	if err != nil {
		if _, ok := err.(*find.NotFoundError); ok {
			// Continue with empty VM slice
		} else {
			return err
		}
	}

	refs := make([]types.ManagedObjectReference, 0, len(vms))
	for _, vm := range vms {
		refs = append(refs, vm.Reference())
	}

	var res infoResult
	var props []string

	if cmd.OutputFlag.JSON {
		props = nil // Load everything
	} else {
		props = []string{"summary"} // Load summary
		if cmd.General {
			props = append(props, "guest.ipAddress")
		}
		if cmd.ExtraConfig {
			props = append(props, "config.extraConfig")
		}
		if cmd.Resources {
			props = append(props, "datastore", "network")
		}
	}

	ctx := context.TODO()
	pc := property.DefaultCollector(c)
	if len(refs) != 0 {
		err = pc.Retrieve(ctx, refs, props, &res.VirtualMachines)
		if err != nil {
			return err
		}
	}

	if cmd.WaitForIP {
		for i, vm := range res.VirtualMachines {
			if vm.Guest == nil || vm.Guest.IpAddress == "" {
				_, err = vms[i].WaitForIP(ctx)
				if err != nil {
					return err
				}
				// Reload virtual machine object
				err = pc.RetrieveOne(ctx, vms[i].Reference(), props, &res.VirtualMachines[i])
				if err != nil {
					return err
				}
			}
		}
	}

	if !cmd.OutputFlag.JSON {
		res.objects = vms
		res.cmd = cmd
		if err = res.collectReferences(pc, ctx); err != nil {
			return err
		}
	}

	return cmd.WriteResult(&res)
}

type infoResult struct {
	VirtualMachines []mo.VirtualMachine
	objects         []*object.VirtualMachine
	entities        map[types.ManagedObjectReference]string
	cmd             *info
}

// collectReferences builds a unique set of MORs to the set of VirtualMachines,
// so we can collect properties in a single call for each reference type {host,datastore,network}.
func (r *infoResult) collectReferences(pc *property.Collector, ctx context.Context) error {
	r.entities = make(map[types.ManagedObjectReference]string) // MOR -> Name map

	var host []mo.HostSystem
	var network []mo.Network
	var dvp []mo.DistributedVirtualPortgroup
	var datastore []mo.Datastore
	// Table to drive inflating refs to their mo.* counterparts (dest)
	// and save() the Name to r.entities w/o using reflection here.
	vrefs := map[string]*struct {
		dest interface{}
		refs []types.ManagedObjectReference
		save func()
	}{
		"host": {
			&host, nil, func() {
				for _, e := range host {
					r.entities[e.Reference()] = e.Name
				}
			},
		},
		"network": {
			&network, nil, func() {
				for _, e := range network {
					r.entities[e.Reference()] = e.Name
				}
			},
		},
		"dvp": {
			&dvp, nil, func() {
				for _, e := range dvp {
					r.entities[e.Reference()] = e.Name
				}
			},
		},
		"datastore": {
			&datastore, nil, func() {
				for _, e := range datastore {
					r.entities[e.Reference()] = e.Name
				}
			},
		},
	}

	xrefs := make(map[types.ManagedObjectReference]bool)
	// Add MOR to vrefs[kind].refs avoiding any duplicates.
	addRef := func(kind string, refs ...types.ManagedObjectReference) {
		for _, ref := range refs {
			if _, exists := xrefs[ref]; exists {
				return
			}
			xrefs[ref] = true
			vref := vrefs[kind]
			vref.refs = append(vref.refs, ref)
		}
	}

	for _, vm := range r.VirtualMachines {
		if r.cmd.General {
			if ref := vm.Summary.Runtime.Host; ref != nil {
				addRef("host", *ref)
			}
		}

		if r.cmd.Resources {
			addRef("datastore", vm.Datastore...)

			for _, net := range vm.Network {
				switch net.Type {
				case "Network":
					addRef("network", net)
				case "DistributedVirtualPortgroup":
					addRef("dvp", net)
				}
			}
		}
	}

	for _, vref := range vrefs {
		if vref.refs == nil {
			continue
		}
		err := pc.Retrieve(ctx, vref.refs, []string{"name"}, vref.dest)
		if err != nil {
			return err
		}
		vref.save()
	}

	return nil
}

func (r *infoResult) entityNames(refs []types.ManagedObjectReference) string {
	var names []string
	for _, ref := range refs {
		names = append(names, r.entities[ref])
	}
	return strings.Join(names, ", ")
}

func (r *infoResult) Write(w io.Writer) error {
	// Maintain order via r.objects as Property collector does not always return results in order.
	objects := make(map[types.ManagedObjectReference]mo.VirtualMachine, len(r.VirtualMachines))
	for _, o := range r.VirtualMachines {
		objects[o.Reference()] = o
	}

	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	for _, o := range r.objects {
		vm := objects[o.Reference()]
		s := vm.Summary

		fmt.Fprintf(tw, "Name:\t%s\n", s.Config.Name)

		if r.cmd.General {
			hostName := "<unavailable>"

			if href := vm.Summary.Runtime.Host; href != nil {
				if name, ok := r.entities[*href]; ok {
					hostName = name
				}
			}

			fmt.Fprintf(tw, "  Path:\t%s\n", o.InventoryPath)
			fmt.Fprintf(tw, "  UUID:\t%s\n", s.Config.Uuid)
			fmt.Fprintf(tw, "  Guest name:\t%s\n", s.Config.GuestFullName)
			fmt.Fprintf(tw, "  Memory:\t%dMB\n", s.Config.MemorySizeMB)
			fmt.Fprintf(tw, "  CPU:\t%d vCPU(s)\n", s.Config.NumCpu)
			fmt.Fprintf(tw, "  Power state:\t%s\n", s.Runtime.PowerState)
			fmt.Fprintf(tw, "  Boot time:\t%s\n", s.Runtime.BootTime)
			fmt.Fprintf(tw, "  IP address:\t%s\n", s.Guest.IpAddress)
			fmt.Fprintf(tw, "  Host:\t%s\n", hostName)
		}

		if r.cmd.Resources {
			fmt.Fprintf(tw, "  CPU usage:\t%dMHz\n", s.QuickStats.OverallCpuUsage)
			fmt.Fprintf(tw, "  Host memory usage:\t%dMB\n", s.QuickStats.HostMemoryUsage)
			fmt.Fprintf(tw, "  Guest memory usage:\t%dMB\n", s.QuickStats.GuestMemoryUsage)
			fmt.Fprintf(tw, "  Storage uncommitted:\t%s\n", units.ByteSize(s.Storage.Uncommitted))
			fmt.Fprintf(tw, "  Storage committed:\t%s\n", units.ByteSize(s.Storage.Committed))
			fmt.Fprintf(tw, "  Storage unshared:\t%s\n", units.ByteSize(s.Storage.Unshared))
			fmt.Fprintf(tw, "  Storage:\t%s\n", r.entityNames(vm.Datastore))
			fmt.Fprintf(tw, "  Network:\t%s\n", r.entityNames(vm.Network))
		}

		if r.cmd.ExtraConfig {
			fmt.Fprintf(tw, "  ExtraConfig:\n")
			for _, v := range vm.Config.ExtraConfig {
				fmt.Fprintf(tw, "    %s:\t%s\n", v.GetOptionValue().Key, v.GetOptionValue().Value)
			}
		}
	}

	return tw.Flush()
}
