// Copyright 2022 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// GENERATED BY gen_go_data.go
// gen_go_data -package beta -var YAML_instance blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/instance.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/instance.yaml
var YAML_instance = []byte("info:\n  title: Compute/Instance\n  description: The Compute Instance resource\n  x-dcl-struct-name: Instance\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a Instance\n    parameters:\n    - name: Instance\n      required: true\n      description: A full instance of a Instance\n  apply:\n    description: The function used to apply information about a Instance\n    parameters:\n    - name: Instance\n      required: true\n      description: A full instance of a Instance\n  delete:\n    description: The function used to delete a Instance\n    parameters:\n    - name: Instance\n      required: true\n      description: A full instance of a Instance\n  deleteAll:\n    description: The function used to delete all Instance\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: zone\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Instance\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: zone\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Instance:\n      title: Instance\n      x-dcl-id: projects/{{project}}/zones/{{zone}}/instances/{{name}}\n      x-dcl-locations:\n      - zone\n      x-dcl-uses-state-hint: true\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-create: true\n      x-dcl-has-iam: true\n      type: object\n      required:\n      - zone\n      - project\n      properties:\n        canIPForward:\n          type: boolean\n          x-dcl-go-name: CanIPForward\n          description: Allows this instance to send and receive packets with non-matching\n            destination or source IPs. This is required if you plan to use this instance\n            to forward routes.\n          x-kubernetes-immutable: true\n        cpuPlatform:\n          type: string\n          x-dcl-go-name: CpuPlatform\n          readOnly: true\n          description: The CPU platform used by this instance.\n          x-kubernetes-immutable: true\n        creationTimestamp:\n          type: string\n          x-dcl-go-name: CreationTimestamp\n          readOnly: true\n          description: Creation timestamp in RFC3339 text format.\n          x-kubernetes-immutable: true\n        deletionProtection:\n          type: boolean\n          x-dcl-go-name: DeletionProtection\n          description: Whether the resource should be protected against deletion.\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description of this resource.\n          x-kubernetes-immutable: true\n        disks:\n          type: array\n          x-dcl-go-name: Disks\n          description: An array of disks that are associated with the instances that\n            are created from this template.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: InstanceDisks\n            properties:\n              autoDelete:\n                type: boolean\n                x-dcl-go-name: AutoDelete\n                description: 'Specifies whether the disk will be auto-deleted when\n                  the instance is deleted (but not when the disk is detached from\n                  the instance).  Tip: Disks should be set to autoDelete=true so that\n                  leftover disks are not left behind on machine deletion.'\n                x-kubernetes-immutable: true\n              boot:\n                type: boolean\n                x-dcl-go-name: Boot\n                description: Indicates that this is a boot disk. The virtual machine\n                  will use the first partition of the disk for its root filesystem.\n                x-kubernetes-immutable: true\n              deviceName:\n                type: string\n                x-dcl-go-name: DeviceName\n                description: Specifies a unique device name of your choice that is\n                  reflected into the /dev/disk/by-id/google-* tree of a Linux operating\n                  system running within the instance. This name can be used to reference\n                  the device for mounting, resizing, and so on, from within the instance.\n                x-kubernetes-immutable: true\n              diskEncryptionKey:\n                type: object\n                x-dcl-go-name: DiskEncryptionKey\n                x-dcl-go-type: InstanceDisksDiskEncryptionKey\n                description: Encrypts or decrypts a disk using a customer-supplied\n                  encryption key.\n                x-kubernetes-immutable: true\n                properties:\n                  rawKey:\n                    type: string\n                    x-dcl-go-name: RawKey\n                    description: Specifies a 256-bit customer-supplied encryption\n                      key, encoded in RFC 4648 base64 to either encrypt or decrypt\n                      this resource.\n                    x-kubernetes-immutable: true\n                  rsaEncryptedKey:\n                    type: string\n                    x-dcl-go-name: RsaEncryptedKey\n                    description: Specifies an RFC 4648 base64 encoded, RSA-wrapped\n                      2048-bit customer-supplied encryption key to either encrypt\n                      or decrypt this resource.\n                    x-kubernetes-immutable: true\n                  sha256:\n                    type: string\n                    x-dcl-go-name: Sha256\n                    readOnly: true\n                    description: The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied\n                      encryption key that protects this resource.\n                    x-kubernetes-immutable: true\n              index:\n                type: integer\n                format: int64\n                x-dcl-go-name: Index\n                description: Assigns a zero-based index to this disk, where 0 is reserved\n                  for the boot disk. For example, if you have many disks attached\n                  to an instance, each disk would have a unique index number. If not\n                  specified, the server will choose an appropriate value.\n                x-kubernetes-immutable: true\n              initializeParams:\n                type: object\n                x-dcl-go-name: InitializeParams\n                x-dcl-go-type: InstanceDisksInitializeParams\n                description: Specifies the parameters for a new disk that will be\n                  created alongside the new instance. Use initialization parameters\n                  to create boot disks or local SSDs attached to the new instance.\n                x-kubernetes-immutable: true\n                x-dcl-mutable-unreadable: true\n                properties:\n                  diskName:\n                    type: string\n                    x-dcl-go-name: DiskName\n                    description: Specifies the disk name. If not specified, the default\n                      is to use the name of the instance.\n                    x-kubernetes-immutable: true\n                  diskSizeGb:\n                    type: integer\n                    format: int64\n                    x-dcl-go-name: DiskSizeGb\n                    description: Specifies the size of the disk in base-2 GB.\n                    x-kubernetes-immutable: true\n                  diskType:\n                    type: string\n                    x-dcl-go-name: DiskType\n                    description: Reference to a disk type. Specifies the disk type\n                      to use to create the instance. If not specified, the default\n                      is pd-standard.\n                    x-kubernetes-immutable: true\n                    x-dcl-references:\n                    - resource: Compute/DiskType\n                      field: name\n                  sourceImage:\n                    type: string\n                    x-dcl-go-name: SourceImage\n                    description: The source image to create this disk. When creating\n                      a new instance, one of initializeParams.sourceImage or disks.source\n                      is required.  To create a disk with one of the public operating\n                      system images, specify the image by its family name.\n                    x-kubernetes-immutable: true\n                  sourceImageEncryptionKey:\n                    type: object\n                    x-dcl-go-name: SourceImageEncryptionKey\n                    x-dcl-go-type: InstanceDisksInitializeParamsSourceImageEncryptionKey\n                    description: The customer-supplied encryption key of the source\n                      image. Required if the source image is protected by a customer-supplied\n                      encryption key.  Instance templates do not store customer-supplied\n                      encryption keys, so you cannot create disks for instances in\n                      a managed instance group if the source images are encrypted\n                      with your own keys.\n                    x-kubernetes-immutable: true\n                    properties:\n                      rawKey:\n                        type: string\n                        x-dcl-go-name: RawKey\n                        description: Specifies a 256-bit customer-supplied encryption\n                          key, encoded in RFC 4648 base64 to either encrypt or decrypt\n                          this resource.\n                        x-kubernetes-immutable: true\n                      sha256:\n                        type: string\n                        x-dcl-go-name: Sha256\n                        readOnly: true\n                        description: The RFC 4648 base64 encoded SHA-256 hash of the\n                          customer-supplied encryption key that protects this resource.\n                        x-kubernetes-immutable: true\n              interface:\n                type: string\n                x-dcl-go-name: Interface\n                x-dcl-go-type: InstanceDisksInterfaceEnum\n                description: Specifies the disk interface to use for attaching this\n                  disk, which is either SCSI or NVME. The default is SCSI. Persistent\n                  disks must always use SCSI and the request will fail if you attempt\n                  to attach a persistent disk in any other format than SCSI.\n                x-kubernetes-immutable: true\n                enum:\n                - SCSI\n                - NVME\n              mode:\n                type: string\n                x-dcl-go-name: Mode\n                x-dcl-go-type: InstanceDisksModeEnum\n                description: The mode in which to attach this disk, either READ_WRITE\n                  or READ_ONLY. If not specified, the default is to attach the disk\n                  in READ_WRITE mode.\n                x-kubernetes-immutable: true\n                enum:\n                - READ_WRITE\n                - READ_ONLY\n              source:\n                type: string\n                x-dcl-go-name: Source\n                description: Reference to a disk. When creating a new instance, one\n                  of initializeParams.sourceImage or disks.source is required.  If\n                  desired, you can also attach existing non-root persistent disks\n                  using this property. This field is only applicable for persistent\n                  disks.\n                x-kubernetes-immutable: true\n                x-dcl-references:\n                - resource: Compute/Disk\n                  field: selfLink\n              type:\n                type: string\n                x-dcl-go-name: Type\n                x-dcl-go-type: InstanceDisksTypeEnum\n                description: Specifies the type of the disk, either SCRATCH or PERSISTENT.\n                  If not specified, the default is PERSISTENT.\n                x-kubernetes-immutable: true\n                enum:\n                - SCRATCH\n                - PERSISTENT\n          x-dcl-mutable-unreadable: true\n        guestAccelerators:\n          type: array\n          x-dcl-go-name: GuestAccelerators\n          description: List of the type and count of accelerator cards attached to\n            the instance\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: InstanceGuestAccelerators\n            properties:\n              acceleratorCount:\n                type: integer\n                format: int64\n                x-dcl-go-name: AcceleratorCount\n                description: The number of the guest accelerator cards exposed to\n                  this instance.\n                x-kubernetes-immutable: true\n              acceleratorType:\n                type: string\n                x-dcl-go-name: AcceleratorType\n                description: Full or partial URL of the accelerator type resource\n                  to expose to this instance.\n                x-kubernetes-immutable: true\n        hostname:\n          type: string\n          x-dcl-go-name: Hostname\n          description: The hostname of the instance to be created. The specified hostname\n            must be RFC1035 compliant. If hostname is not specified, the default hostname\n            is [INSTANCE_NAME].c.[PROJECT_ID].internal when using the global DNS,\n            and [INSTANCE_NAME].[ZONE].c.[PROJECT_ID].internal when using zonal DNS.\n          x-kubernetes-immutable: true\n        id:\n          type: string\n          x-dcl-go-name: Id\n          readOnly: true\n          description: The unique identifier for the resource. This identifier is\n            defined by the server.\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Labels to apply to this instance.  A list of key->value pairs.\n        machineType:\n          type: string\n          x-dcl-go-name: MachineType\n          description: A reference to a machine type which defines VM kind.\n          x-dcl-forward-slash-allowed: true\n          x-dcl-references:\n          - resource: Compute/MachineType\n            field: name\n        metadata:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Metadata\n          description: The metadata key/value pairs to assign to instances that are\n            created from this template. These pairs can consist of custom metadata\n            or predefined keys.\n        minCpuPlatform:\n          type: string\n          x-dcl-go-name: MinCpuPlatform\n          description: Specifies a minimum CPU platform for the VM instance. Applicable\n            values are the friendly names of CPU platforms\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of the resource, provided by the client when initially\n            creating the resource. The resource name must be 1-63 characters long,\n            and comply with RFC1035. Specifically, the name must be 1-63 characters\n            long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which\n            means the first character must be a lowercase letter, and all following\n            characters must be a dash, lowercase letter, or digit, except the last\n            character, which cannot be a dash.\n          x-kubernetes-immutable: true\n        networkInterfaces:\n          type: array\n          x-dcl-go-name: NetworkInterfaces\n          description: An array of configurations for this interface. This specifies\n            how this interface is configured to interact with other network services,\n            such as connecting to the internet. Only one network interface is supported\n            per instance.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: InstanceNetworkInterfaces\n            properties:\n              accessConfigs:\n                type: array\n                x-dcl-go-name: AccessConfigs\n                description: An array of configurations for this interface. Currently,\n                  only one access config, ONE_TO_ONE_NAT, is supported. If there are\n                  no accessConfigs specified, then this instance will have no external\n                  internet access.\n                x-kubernetes-immutable: true\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: object\n                  x-dcl-go-type: InstanceNetworkInterfacesAccessConfigs\n                  required:\n                  - name\n                  - type\n                  properties:\n                    externalIPv6:\n                      type: string\n                      x-dcl-go-name: ExternalIPv6\n                      readOnly: true\n                      description: The first IPv6 address of the external IPv6 range\n                        associated with this instance, prefix length is stored in\n                        externalIpv6PrefixLength in ipv6AccessConfig. The field is\n                        output only, an IPv6 address from a subnetwork associated\n                        with the instance will be allocated dynamically.\n                      x-kubernetes-immutable: true\n                    externalIPv6PrefixLength:\n                      type: string\n                      x-dcl-go-name: ExternalIPv6PrefixLength\n                      readOnly: true\n                      description: The prefix length of the external IPv6 range.\n                      x-kubernetes-immutable: true\n                    name:\n                      type: string\n                      x-dcl-go-name: Name\n                      description: The name of this access configuration. The default\n                        and recommended name is External NAT but you can use any arbitrary\n                        string you would like. For example, My external IP or Network\n                        Access.\n                      x-kubernetes-immutable: true\n                    natIP:\n                      type: string\n                      x-dcl-go-name: NatIP\n                      description: Reference to an address. An external IP address\n                        associated with this instance. Specify an unused static external\n                        IP address available to the project or leave this field undefined\n                        to use an IP from a shared ephemeral IP address pool. If you\n                        specify a static external IP address, it must live in the\n                        same region as the zone of the instance.\n                      x-kubernetes-immutable: true\n                      x-dcl-references:\n                      - resource: Compute/Address\n                        field: selfLink\n                    networkTier:\n                      type: string\n                      x-dcl-go-name: NetworkTier\n                      x-dcl-go-type: InstanceNetworkInterfacesAccessConfigsNetworkTierEnum\n                      description: 'This signifies the networking tier used for configuring\n                        this access configuration and can only take the following\n                        values: PREMIUM, STANDARD. If an AccessConfig is specified\n                        without a valid external IP address, an ephemeral IP will\n                        be created with this networkTier. If an AccessConfig with\n                        a valid external IP address is specified, it must match that\n                        of the networkTier associated with the Address resource owning\n                        that IP.'\n                      x-kubernetes-immutable: true\n                      enum:\n                      - PREMIUM\n                      - STANDARD\n                    publicPtrDomainName:\n                      type: string\n                      x-dcl-go-name: PublicPtrDomainName\n                      description: The DNS domain name for the public PTR record.\n                        You can set this field only if the setPublicPtr field is enabled.\n                      x-kubernetes-immutable: true\n                    setPublicPtr:\n                      type: boolean\n                      x-dcl-go-name: SetPublicPtr\n                      description: Specifies whether a public DNS 'PTR' record should\n                        be created to map the external IP address of the instance\n                        to a DNS domain name.\n                      x-kubernetes-immutable: true\n                    type:\n                      type: string\n                      x-dcl-go-name: Type\n                      x-dcl-go-type: InstanceNetworkInterfacesAccessConfigsTypeEnum\n                      description: The type of configuration. The default and only\n                        option is ONE_TO_ONE_NAT.\n                      x-kubernetes-immutable: true\n                      enum:\n                      - ONE_TO_ONE_NAT\n              aliasIPRanges:\n                type: array\n                x-dcl-go-name: AliasIPRanges\n                description: An array of alias IP ranges for this network interface.\n                  Can only be specified for network interfaces on subnet-mode networks.\n                x-kubernetes-immutable: true\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: object\n                  x-dcl-go-type: InstanceNetworkInterfacesAliasIPRanges\n                  properties:\n                    ipCidrRange:\n                      type: string\n                      x-dcl-go-name: IPCidrRange\n                      description: The IP CIDR range represented by this alias IP\n                        range. This IP CIDR range must belong to the specified subnetwork\n                        and cannot contain IP addresses reserved by system or used\n                        by other network interfaces. This range may be a single IP\n                        address (e.g. 10.2.3.4), a netmask (e.g. /24) or a CIDR format\n                        string (e.g. 10.1.2.0/24).\n                      x-kubernetes-immutable: true\n                    subnetworkRangeName:\n                      type: string\n                      x-dcl-go-name: SubnetworkRangeName\n                      description: Optional subnetwork secondary range name specifying\n                        the secondary range from which to allocate the IP CIDR range\n                        for this alias IP range. If left unspecified, the primary\n                        range of the subnetwork will be used.\n                      x-kubernetes-immutable: true\n              ipv6AccessConfigs:\n                type: array\n                x-dcl-go-name: IPv6AccessConfigs\n                description: An array of IPv6 access configurations for this interface.\n                  Currently, only one IPv6 access config, DIRECT_IPV6, is supported.\n                  If there is no ipv6AccessConfig specified, then this instance will\n                  have no external IPv6 Internet access.\n                x-kubernetes-immutable: true\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: object\n                  x-dcl-go-type: InstanceNetworkInterfacesIPv6AccessConfigs\n                  required:\n                  - name\n                  - type\n                  properties:\n                    externalIPv6:\n                      type: string\n                      x-dcl-go-name: ExternalIPv6\n                      readOnly: true\n                      description: The first IPv6 address of the external IPv6 range\n                        associated with this instance, prefix length is stored in\n                        externalIpv6PrefixLength in ipv6AccessConfig. The field is\n                        output only, an IPv6 address from a subnetwork associated\n                        with the instance will be allocated dynamically.\n                      x-kubernetes-immutable: true\n                    externalIPv6PrefixLength:\n                      type: string\n                      x-dcl-go-name: ExternalIPv6PrefixLength\n                      readOnly: true\n                      description: The prefix length of the external IPv6 range.\n                      x-kubernetes-immutable: true\n                    name:\n                      type: string\n                      x-dcl-go-name: Name\n                      description: The name of this access configuration. The default\n                        and recommended name is External NAT but you can use any arbitrary\n                        string you would like. For example, My external IP or Network\n                        Access.\n                      x-kubernetes-immutable: true\n                    natIP:\n                      type: string\n                      x-dcl-go-name: NatIP\n                      description: Reference to an address. An external IP address\n                        associated with this instance. Specify an unused static external\n                        IP address available to the project or leave this field undefined\n                        to use an IP from a shared ephemeral IP address pool. If you\n                        specify a static external IP address, it must live in the\n                        same region as the zone of the instance.\n                      x-kubernetes-immutable: true\n                      x-dcl-references:\n                      - resource: Compute/Address\n                        field: selfLink\n                    networkTier:\n                      type: string\n                      x-dcl-go-name: NetworkTier\n                      x-dcl-go-type: InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum\n                      description: 'This signifies the networking tier used for configuring\n                        this access configuration and can only take the following\n                        values: PREMIUM, STANDARD. If an AccessConfig is specified\n                        without a valid external IP address, an ephemeral IP will\n                        be created with this networkTier. If an AccessConfig with\n                        a valid external IP address is specified, it must match that\n                        of the networkTier associated with the Address resource owning\n                        that IP.'\n                      x-kubernetes-immutable: true\n                      enum:\n                      - PREMIUM\n                      - STANDARD\n                    publicPtrDomainName:\n                      type: string\n                      x-dcl-go-name: PublicPtrDomainName\n                      description: The DNS domain name for the public PTR record.\n                        You can set this field only if the setPublicPtr field is enabled.\n                      x-kubernetes-immutable: true\n                    setPublicPtr:\n                      type: boolean\n                      x-dcl-go-name: SetPublicPtr\n                      description: Specifies whether a public DNS 'PTR' record should\n                        be created to map the external IP address of the instance\n                        to a DNS domain name.\n                      x-kubernetes-immutable: true\n                    type:\n                      type: string\n                      x-dcl-go-name: Type\n                      x-dcl-go-type: InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum\n                      description: The type of configuration. The default and only\n                        option is ONE_TO_ONE_NAT.\n                      x-kubernetes-immutable: true\n                      enum:\n                      - ONE_TO_ONE_NAT\n              name:\n                type: string\n                x-dcl-go-name: Name\n                readOnly: true\n                description: The name of the network interface, generated by the server.\n                  For network devices, these are eth0, eth1, etc\n                x-kubernetes-immutable: true\n              network:\n                type: string\n                x-dcl-go-name: Network\n                description: Specifies the title of an existing network.  When creating\n                  an instance, if neither the network nor the subnetwork is specified,\n                  the default network global/networks/default is used; if the network\n                  is not specified but the subnetwork is specified, the network is\n                  inferred.\n                x-kubernetes-immutable: true\n                x-dcl-references:\n                - resource: Compute/Network\n                  field: name\n              networkIP:\n                type: string\n                x-dcl-go-name: NetworkIP\n                description: An IPv4 internal network address to assign to the instance\n                  for this network interface. If not specified by the user, an unused\n                  internal IP is assigned by the system.\n                x-kubernetes-immutable: true\n              subnetwork:\n                type: string\n                x-dcl-go-name: Subnetwork\n                description: Reference to a VPC network. If the network resource is\n                  in legacy mode, do not provide this property.  If the network is\n                  in auto subnet mode, providing the subnetwork is optional. If the\n                  network is in custom subnet mode, then this field should be specified.\n                x-kubernetes-immutable: true\n                x-dcl-references:\n                - resource: Compute/Subnetwork\n                  field: name\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project id of the resource.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        scheduling:\n          type: object\n          x-dcl-go-name: Scheduling\n          x-dcl-go-type: InstanceScheduling\n          description: Sets the scheduling options for this instance.\n          x-kubernetes-immutable: true\n          properties:\n            automaticRestart:\n              type: boolean\n              x-dcl-go-name: AutomaticRestart\n              description: Specifies whether the instance should be automatically\n                restarted if it is terminated by Compute Engine (not terminated by\n                a user). You can only set the automatic restart option for standard\n                instances. Preemptible instances cannot be automatically restarted.\n              x-kubernetes-immutable: true\n            onHostMaintenance:\n              type: string\n              x-dcl-go-name: OnHostMaintenance\n              description: Defines the maintenance behavior for this instance. For\n                standard instances, the default behavior is MIGRATE. For preemptible\n                instances, the default and only possible behavior is TERMINATE. For\n                more information, see Setting Instance Scheduling Options.\n              x-kubernetes-immutable: true\n            preemptible:\n              type: boolean\n              x-dcl-go-name: Preemptible\n              description: Defines whether the instance is preemptible. This can only\n                be set during instance creation, it cannot be set or changed after\n                the instance has been created.\n              x-kubernetes-immutable: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: The self link of the instance\n          x-kubernetes-immutable: true\n        serviceAccounts:\n          type: array\n          x-dcl-go-name: ServiceAccounts\n          description: A list of service accounts, with their specified scopes, authorized\n            for this instance. Only one service account per VM instance is supported.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: InstanceServiceAccounts\n            properties:\n              email:\n                type: string\n                x-dcl-go-name: Email\n                description: Email address of the service account.\n                x-kubernetes-immutable: true\n              scopes:\n                type: array\n                x-dcl-go-name: Scopes\n                description: The list of scopes to be made available for this service\n                  account.\n                x-kubernetes-immutable: true\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: string\n                  x-dcl-go-type: string\n        shieldedInstanceConfig:\n          type: object\n          x-dcl-go-name: ShieldedInstanceConfig\n          x-dcl-go-type: InstanceShieldedInstanceConfig\n          description: Configuration for various parameters related to shielded instances.\n          properties:\n            enableIntegrityMonitoring:\n              type: boolean\n              x-dcl-go-name: EnableIntegrityMonitoring\n              description: Defines whether the instance has integrity monitoring enabled.\n            enableSecureBoot:\n              type: boolean\n              x-dcl-go-name: EnableSecureBoot\n              description: Defines whether the instance has Secure Boot enabled.\n            enableVtpm:\n              type: boolean\n              x-dcl-go-name: EnableVtpm\n              description: Defines whether the instance has the vTPM enabled\n        status:\n          type: string\n          x-dcl-go-name: Status\n          x-dcl-go-type: InstanceStatusEnum\n          description: 'The status of the instance. One of the following values: PROVISIONING,\n            STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, and TERMINATED.  As\n            a user, use RUNNING to keep a machine \"on\" and TERMINATED to turn a machine\n            off'\n          x-kubernetes-immutable: true\n          enum:\n          - PROVISIONING\n          - STAGING\n          - RUNNING\n          - STOPPING\n          - SUSPENDING\n          - SUSPENDED\n          - TERMINATED\n        statusMessage:\n          type: string\n          x-dcl-go-name: StatusMessage\n          readOnly: true\n          description: An optional, human-readable explanation of the status.\n          x-kubernetes-immutable: true\n        tags:\n          type: array\n          x-dcl-go-name: Tags\n          description: A list of tags to apply to this instance. Tags are used to\n            identify valid sources or targets for network firewalls and are specified\n            by the client during instance creation. Each tag within the list must\n            comply with RFC1035.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        zone:\n          type: string\n          x-dcl-go-name: Zone\n          description: A reference to the zone where the machine resides.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Compute/Zone\n            field: name\n            parent: true\n")

// 35498 bytes
// MD5: 2a1322ba1e52f1e0ab72c70d84d49914
