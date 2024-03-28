/*
Copyright 2022 The Kubernetes Authors.

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

package networking

import (
	infrav1 "github.com/easystack/cluster-api-provider-openstack/api/v1alpha6"
)

var defaultRules = []infrav1.SecurityGroupRule{
	{
		Direction:      "egress",
		Description:    "Full open",
		EtherType:      "IPv4",
		PortRangeMin:   0,
		PortRangeMax:   0,
		Protocol:       "",
		RemoteIPPrefix: "",
	},
	{
		Direction:      "egress",
		Description:    "Full open",
		EtherType:      "IPv6",
		PortRangeMin:   0,
		PortRangeMax:   0,
		Protocol:       "",
		RemoteIPPrefix: "",
	},
}

// Permit traffic for cadvisor
func GetSGControlPlaneForCadvisor(remoteGroupIDSelf, secWorkerGroupID string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:   "cadvisor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  4194,
			PortRangeMax:  4194,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "cadvisor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  4194,
			PortRangeMax:  4194,
			Protocol:      "tcp",
			RemoteGroupID: secWorkerGroupID,
		},
	}
}

func GetSGWorkForCadvisor(remoteGroupIDSelf, secControlPlaneGroupID string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:   "cadvisor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  4194,
			PortRangeMax:  4194,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "cadvisor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  4194,
			PortRangeMax:  4194,
			Protocol:      "tcp",
			RemoteGroupID: secControlPlaneGroupID,
		},
	}
}

// Permit traffic for coredns
func GetSGControlPlaneForCOREDNS(remoteGroupIDSelf, secWorkerGroupID string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:  "coredns",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 9153,
			PortRangeMax: 9153,
			Protocol:     "tcp",
		},
		{
			Description:  "coredns",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 9253,
			PortRangeMax: 9253,
			Protocol:     "tcp",
		},
		{
			Description:  "coredns",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 53,
			PortRangeMax: 53,
			Protocol:     "tcp",
		},
		{
			Description:  "coredns",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 53,
			PortRangeMax: 53,
			Protocol:     "udp",
		},
	}
}

func GetSGWorkForCOREDNS(remoteGroupIDSelf, secControlPlaneGroupID string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:  "coredns",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 9153,
			PortRangeMax: 9153,
			Protocol:     "tcp",
		},
		{
			Description:  "coredns",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 9253,
			PortRangeMax: 9253,
			Protocol:     "tcp",
		},
		{
			Description:  "coredns",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 53,
			PortRangeMax: 53,
			Protocol:     "tcp",
		},
		{
			Description:  "coredns",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 53,
			PortRangeMax: 53,
			Protocol:     "udp",
		},
	}
}

// Permit traffic for prometheus
func GetSGControlPlaneForPrometheus(remoteGroupIDSelf, secWorkerGroupID string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9100,
			PortRangeMax:  9100,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9100,
			PortRangeMax:  9100,
			Protocol:      "tcp",
			RemoteGroupID: secWorkerGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9200,
			PortRangeMax:  9200,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9200,
			PortRangeMax:  9200,
			Protocol:      "tcp",
			RemoteGroupID: secWorkerGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9300,
			PortRangeMax:  9300,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9300,
			PortRangeMax:  9300,
			Protocol:      "tcp",
			RemoteGroupID: secWorkerGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9090,
			PortRangeMax:  9090,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9090,
			PortRangeMax:  9090,
			Protocol:      "tcp",
			RemoteGroupID: secWorkerGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9091,
			PortRangeMax:  9091,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9091,
			PortRangeMax:  9091,
			Protocol:      "tcp",
			RemoteGroupID: secWorkerGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9093,
			PortRangeMax:  9093,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9093,
			PortRangeMax:  9093,
			Protocol:      "tcp",
			RemoteGroupID: secWorkerGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9094,
			PortRangeMax:  9094,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9094,
			PortRangeMax:  9094,
			Protocol:      "tcp",
			RemoteGroupID: secWorkerGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9094,
			PortRangeMax:  9094,
			Protocol:      "udp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9094,
			PortRangeMax:  9094,
			Protocol:      "udp",
			RemoteGroupID: secWorkerGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  30300,
			PortRangeMax:  30300,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  30300,
			PortRangeMax:  30300,
			Protocol:      "tcp",
			RemoteGroupID: secWorkerGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  3000,
			PortRangeMax:  3000,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  3000,
			PortRangeMax:  3000,
			Protocol:      "tcp",
			RemoteGroupID: secWorkerGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  10901,
			PortRangeMax:  10901,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  10901,
			PortRangeMax:  10901,
			Protocol:      "tcp",
			RemoteGroupID: secWorkerGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  10902,
			PortRangeMax:  10902,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  10902,
			PortRangeMax:  10902,
			Protocol:      "tcp",
			RemoteGroupID: secWorkerGroupID,
		},
	}
}

func GetSGWorkForPrometheus(remoteGroupIDSelf, secControlPlaneGroupID string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9100,
			PortRangeMax:  9100,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9100,
			PortRangeMax:  9100,
			Protocol:      "tcp",
			RemoteGroupID: secControlPlaneGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9200,
			PortRangeMax:  9200,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9200,
			PortRangeMax:  9200,
			Protocol:      "tcp",
			RemoteGroupID: secControlPlaneGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9300,
			PortRangeMax:  9300,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9300,
			PortRangeMax:  9300,
			Protocol:      "tcp",
			RemoteGroupID: secControlPlaneGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9090,
			PortRangeMax:  9090,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9090,
			PortRangeMax:  9090,
			Protocol:      "tcp",
			RemoteGroupID: secControlPlaneGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9091,
			PortRangeMax:  9091,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9091,
			PortRangeMax:  9091,
			Protocol:      "tcp",
			RemoteGroupID: secControlPlaneGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9093,
			PortRangeMax:  9093,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9093,
			PortRangeMax:  9093,
			Protocol:      "tcp",
			RemoteGroupID: secControlPlaneGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9094,
			PortRangeMax:  9094,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9094,
			PortRangeMax:  9094,
			Protocol:      "tcp",
			RemoteGroupID: secControlPlaneGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9094,
			PortRangeMax:  9094,
			Protocol:      "udp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  9094,
			PortRangeMax:  9094,
			Protocol:      "udp",
			RemoteGroupID: secControlPlaneGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  30300,
			PortRangeMax:  30300,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  30300,
			PortRangeMax:  30300,
			Protocol:      "tcp",
			RemoteGroupID: secControlPlaneGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  3000,
			PortRangeMax:  3000,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  3000,
			PortRangeMax:  3000,
			Protocol:      "tcp",
			RemoteGroupID: secControlPlaneGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  10901,
			PortRangeMax:  10901,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  10901,
			PortRangeMax:  10901,
			Protocol:      "tcp",
			RemoteGroupID: secControlPlaneGroupID,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  10902,
			PortRangeMax:  10902,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "prometheus-monitor",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  10902,
			PortRangeMax:  10902,
			Protocol:      "tcp",
			RemoteGroupID: secControlPlaneGroupID,
		},
	}
}

// Permit traffic for etcd, kubelet , kube-scheduler,kube-controller-manager
func getSGControlPlaneCommon(remoteGroupIDSelf, secWorkerGroupID string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:  "Etcd",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 2379,
			PortRangeMax: 2380,
			Protocol:     "tcp",
		},
		{
			// kubeadm says this is needed
			Description:   "Kubelet API",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  10240,
			PortRangeMax:  10260,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			// This is needed to support metrics-server deployments
			Description:   "Kubelet API",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  10240,
			PortRangeMax:  10260,
			Protocol:      "tcp",
			RemoteGroupID: secWorkerGroupID,
		},
	}
}

// Permit traffic for calico.
func getSGControlPlaneCalico(remoteGroupIDSelf, secWorkerGroupID string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:   "BGP (calico)",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  179,
			PortRangeMax:  179,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "BGP (calico)",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  179,
			PortRangeMax:  179,
			Protocol:      "tcp",
			RemoteGroupID: secWorkerGroupID,
		},
		{
			Description:   "IP-in-IP (calico)",
			Direction:     "ingress",
			EtherType:     "IPv4",
			Protocol:      "4",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "IP-in-IP (calico)",
			Direction:     "ingress",
			EtherType:     "IPv4",
			Protocol:      "4",
			RemoteGroupID: secWorkerGroupID,
		},
	}
}

// Permit traffic for kubelet.
func getSGWorkerCommon(remoteGroupIDSelf, secControlPlaneGroupID string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			// This is needed to support metrics-server deployments
			Description:   "Kubelet API",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  10250,
			PortRangeMax:  10250,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "Kubelet API",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  10250,
			PortRangeMax:  10250,
			Protocol:      "tcp",
			RemoteGroupID: secControlPlaneGroupID,
		},
	}
}

// Permit traffic for calico.
func getSGWorkerCalico(remoteGroupIDSelf, secControlPlaneGroupID string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:   "BGP (calico)",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  179,
			PortRangeMax:  179,
			Protocol:      "tcp",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "BGP (calico)",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  179,
			PortRangeMax:  179,
			Protocol:      "tcp",
			RemoteGroupID: secControlPlaneGroupID,
		},
		{
			Description:   "IP-in-IP (calico)",
			Direction:     "ingress",
			EtherType:     "IPv4",
			Protocol:      "4",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "IP-in-IP (calico)",
			Direction:     "ingress",
			EtherType:     "IPv4",
			Protocol:      "4",
			RemoteGroupID: secControlPlaneGroupID,
		},
	}
}

// Permit traffic for ssh control plane.
func GetSGControlPlaneSSH(secBastionGroupID string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:   "SSH",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  22,
			PortRangeMax:  22,
			Protocol:      "tcp",
			RemoteGroupID: secBastionGroupID,
		},
	}
}

// Permit traffic for ssh worker.
func GetSGWorkerSSH(secBastionGroupID string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:   "SSH",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  22,
			PortRangeMax:  22,
			Protocol:      "tcp",
			RemoteGroupID: secBastionGroupID,
		},
	}
}

// Allow icmp traffic from control plane.
func GetSGControlPlaneICMP(remoteGroupIDSelf string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description: "ICMP",
			Direction:   "ingress",
			EtherType:   "IPv4",
			Protocol:    "icmp",
		},
	}
}

// Allow icmp traffic from worker.
func GetSGWorkerICMP(remoteGroupIDSelf string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description: "ICMP",
			Direction:   "ingress",
			EtherType:   "IPv4",
			Protocol:    "icmp",
		},
	}
}

// Allow all traffic, including from outside the cluster, to access the ingress API
func GetSGControlPlaneOrWorkIngress() []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:  "Ingress API",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 80,
			PortRangeMax: 80,
			Protocol:     "tcp",
		},
		{
			Description:  "Ingress API",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 443,
			PortRangeMax: 443,
			Protocol:     "tcp",
		},
	}
}

// Permit traffic for flannel.
func GetSGControlPlaneFlannel() []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:  "flannel",
			Direction:    "ingress",
			EtherType:    "IPv4",
			Protocol:     "udp",
			PortRangeMin: 8472,
			PortRangeMax: 8472,
		},
	}
}

func GetSGWorkerFlannel() []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:  "flannel",
			Direction:    "ingress",
			EtherType:    "IPv4",
			Protocol:     "udp",
			PortRangeMin: 8472,
			PortRangeMax: 8472,
		},
	}
}

// Permit traffic for keepalived
func GetSGControlPlaneOrWorkVRRP() []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description: "keepalived",
			Direction:   "ingress",
			EtherType:   "IPv4",
			Protocol:    "vrrp",
		},
	}
}

// Allow all traffic, including from outside the cluster, to access the API.
func GetSGControlPlaneHTTPS() []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:  "Kubernetes API",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 6443,
			PortRangeMax: 6443,
			Protocol:     "tcp",
		},
	}
}

// Allow all traffic, including from outside the cluster, to access the API from 8080
func GetSGControlPlaneHTTP() []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:  "Kubernetes API",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 8080,
			PortRangeMax: 8080,
			Protocol:     "tcp",
		},
	}
}

// Allow all traffic, including from outside the cluster, to access the API from nginx
func GetSGControlPlaneHTTPSNGINX() []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:  "Kubernetes API",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 8443,
			PortRangeMax: 8443,
			Protocol:     "tcp",
		},
	}
}

// Allow all traffic, including from outside the cluster, to access node port services.
func GetSGControlPlaneNodePort() []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:  "Node Port Services",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 30000,
			PortRangeMax: 32767,
			Protocol:     "tcp",
		},
		{
			Description:  "Node Port Services",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 30000,
			PortRangeMax: 32767,
			Protocol:     "udp",
		},
	}
}

// Allow all traffic, including from outside the cluster, to access node port services.
func GetSGWorkerNodePort() []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:  "Node Port Services",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 30000,
			PortRangeMax: 32767,
			Protocol:     "tcp",
		},
		{
			Description:  "Node Port Services",
			Direction:    "ingress",
			EtherType:    "IPv4",
			PortRangeMin: 30000,
			PortRangeMax: 32767,
			Protocol:     "udp",
		},
	}
}

// Permit all ingress from the cluster security groups.
func GetSGControlPlaneAllowAll(remoteGroupIDSelf, secWorkerGroupID string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:   "In-cluster Ingress",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  0,
			PortRangeMax:  0,
			Protocol:      "",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "In-cluster Ingress",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  0,
			PortRangeMax:  0,
			Protocol:      "",
			RemoteGroupID: secWorkerGroupID,
		},
	}
}

// Permit all ingress from the cluster security groups.
func GetSGWorkerAllowAll(remoteGroupIDSelf, secControlPlaneGroupID string) []infrav1.SecurityGroupRule {
	return []infrav1.SecurityGroupRule{
		{
			Description:   "In-cluster Ingress",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  0,
			PortRangeMax:  0,
			Protocol:      "",
			RemoteGroupID: remoteGroupIDSelf,
		},
		{
			Description:   "In-cluster Ingress",
			Direction:     "ingress",
			EtherType:     "IPv4",
			PortRangeMin:  0,
			PortRangeMax:  0,
			Protocol:      "",
			RemoteGroupID: secControlPlaneGroupID,
		},
	}
}

func GetSGControlPlaneGeneral(remoteGroupIDSelf, secWorkerGroupID string) []infrav1.SecurityGroupRule {
	controlPlaneRules := []infrav1.SecurityGroupRule{}
	controlPlaneRules = append(controlPlaneRules, getSGControlPlaneCommon(remoteGroupIDSelf, secWorkerGroupID)...)
	controlPlaneRules = append(controlPlaneRules, getSGControlPlaneCalico(remoteGroupIDSelf, secWorkerGroupID)...)
	return controlPlaneRules
}

func GetSGWorkerGeneral(remoteGroupIDSelf, secControlPlaneGroupID string) []infrav1.SecurityGroupRule {
	workerRules := []infrav1.SecurityGroupRule{}
	workerRules = append(workerRules, getSGWorkerCommon(remoteGroupIDSelf, secControlPlaneGroupID)...)
	workerRules = append(workerRules, getSGWorkerCalico(remoteGroupIDSelf, secControlPlaneGroupID)...)
	return workerRules
}
