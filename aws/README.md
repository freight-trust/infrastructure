
 - version: 4
 - terraform_version: 0.12.24
 - serial: 70
 - lineage: 897dbb43-ded2-a543-5d71-5bd06caa41cd
 - outputs
    
     - azs
        
         - value
            
             - 0: us-west-1a
             - 1: us-west-1c
            
         - type
            
             - 0: list
             - 1: string
            
        
     - bootnode_ip
        
         - value: 3.101.83.95
         - type: string
        
     - monitoring_ip
        
         - value: 18.144.58.233
         - type: string
        
     - nat_public_ips
        
         - value
            
             - 0: 13.52.210.124
            
         - type
            
             - 0: tuple
             - 1
                
                 - 0: string
                
            
        
     - node_ip
        
         - value
            
             - 0: 54.183.133.228
             - 1: 54.67.10.205
             - 2: 54.193.83.95
             - 3: 54.193.114.99
             - 4: 54.183.62.83
            
         - type
            
             - 0: tuple
             - 1
                
                 - 0: string
                 - 1: string
                 - 2: string
                 - 3: string
                 - 4: string
                
            
        
     - private_subnets
        
         - value
            
             - 0: subnet-06e62b6d0c1d8e1b0
             - 1: subnet-003fdf10e9015f50a
             - 2: subnet-0d97d5bad724ebcac
            
         - type
            
             - 0: tuple
             - 1
                
                 - 0: string
                 - 1: string
                 - 2: string
                
            
        
     - public_subnets
        
         - value
            
             - 0: subnet-07df858a84da18c00
             - 1: subnet-0d619244f43aeea5e
             - 2: subnet-0005aeeaff667e9f9
            
         - type
            
             - 0: tuple
             - 1
                
                 - 0: string
                 - 1: string
                 - 2: string
                
            
        
     - vpc_cidr_block
        
         - value: 10.0.0.0/16
         - type: string
        
     - vpc_id
        
         - value: vpc-0881077c57fc53714
         - type: string
        
    
 - resources
    
     - 0
        
         - mode: data
         - type: aws_ami
         - name: ubuntu
         - provider: provider.aws
         - instances
            
             - 0
                
                 - schema_version: 0
                 - attributes
                    
                     - architecture: x86_64
                     - block_device_mappings
                        
                         - 0
                            
                             - device_name: /dev/sda1
                             - ebs
                                
                                 - delete_on_termination: true
                                 - encrypted: false
                                 - iops: 0
                                 - snapshot_id: snap-047a5d224a46c4b97
                                 - volume_size: 8
                                 - volume_type: gp2
                                
                             - no_device: 
                             - virtual_name: 
                            
                         - 1
                            
                             - device_name: /dev/sdb
                             - ebs
                                
                                
                             - no_device: 
                             - virtual_name: ephemeral0
                            
                         - 2
                            
                             - device_name: /dev/sdc
                             - ebs
                                
                                
                             - no_device: 
                             - virtual_name: ephemeral1
                            
                        
                     - creation_date: 2020-06-11T22:05:41.000Z
                     - description: Canonical, Ubuntu, 18.04 LTS, amd64 bionic image build on 2020-06-11
                     - executable_users: null
                     - filter
                        
                         - 0
                            
                             - name: name
                             - values
                                
                                 - 0: ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-*
                                
                            
                        
                     - hypervisor: xen
                     - id: ami-0d705db840ec5f0c5
                     - image_id: ami-0d705db840ec5f0c5
                     - image_location: 099720109477/ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-20200611
                     - image_owner_alias: null
                     - image_type: machine
                     - kernel_id: null
                     - most_recent: true
                     - name: ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-20200611
                     - name_regex: null
                     - owner_id: 099720109477
                     - owners
                        
                         - 0: 099720109477
                        
                     - platform: null
                     - product_codes
                        
                        
                     - public: true
                     - ramdisk_id: null
                     - root_device_name: /dev/sda1
                     - root_device_type: ebs
                     - root_snapshot_id: snap-047a5d224a46c4b97
                     - sriov_net_support: simple
                     - state: available
                     - state_reason
                        
                         - code: UNSET
                         - message: UNSET
                        
                     - tags
                        
                        
                     - virtualization_type: hvm
                    
                
            
        
     - 1
        
         - module: module.vpc
         - mode: data
         - type: aws_iam_policy_document
         - name: flow_log_cloudwatch_assume_role
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 2
        
         - module: module.vpc
         - mode: data
         - type: aws_iam_policy_document
         - name: vpc_flow_log_cloudwatch
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 3
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: access_analyzer
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 4
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: acm_pca
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 5
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: apigw
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 6
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: appmesh_envoy_management
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 7
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: appstream
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 8
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: athena
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 9
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: auto_scaling_plans
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 10
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: cloud_directory
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 11
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: cloudformation
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 12
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: cloudtrail
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 13
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: codebuild
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 14
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: codecommit
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 15
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: codepipeline
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 16
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: config
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 17
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: datasync
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 18
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: dynamodb
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 19
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: ebs
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 20
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: ec2
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 21
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: ec2_autoscaling
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 22
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: ec2messages
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 23
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: ecr_api
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 24
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: ecr_dkr
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 25
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: ecs
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 26
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: ecs_agent
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 27
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: ecs_telemetry
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 28
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: efs
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 29
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: elastic_inference_runtime
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 30
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: elasticbeanstalk
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 31
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: elasticbeanstalk_health
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 32
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: elasticloadbalancing
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 33
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: emr
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 34
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: events
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 35
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: git_codecommit
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 36
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: glue
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 37
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: kinesis_firehose
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 38
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: kinesis_streams
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 39
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: kms
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 40
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: logs
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 41
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: monitoring
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 42
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: qldb_session
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 43
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: rekognition
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 44
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: s3
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 45
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: sagemaker_api
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 46
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: sagemaker_notebook
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 47
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: sagemaker_runtime
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 48
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: secretsmanager
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 49
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: servicecatalog
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 50
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: ses
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 51
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: sms
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 52
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: sns
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 53
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: sqs
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 54
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: ssm
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 55
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: ssmmessages
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 56
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: states
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 57
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: storagegateway
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 58
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: sts
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 59
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: transfer
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 60
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: transferserver
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 61
        
         - module: module.vpc
         - mode: data
         - type: aws_vpc_endpoint_service
         - name: workspaces
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 62
        
         - module: module.vpc
         - mode: managed
         - type: aws_cloudwatch_log_group
         - name: flow_log
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 63
        
         - module: module.vpc
         - mode: managed
         - type: aws_customer_gateway
         - name: this
         - each: map
         - provider: provider.aws
         - instances
            
            
        
     - 64
        
         - module: module.vpc
         - mode: managed
         - type: aws_db_subnet_group
         - name: database
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 65
        
         - module: module.vpc
         - mode: managed
         - type: aws_default_network_acl
         - name: this
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 66
        
         - module: module.vpc
         - mode: managed
         - type: aws_default_vpc
         - name: this
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 67
        
         - module: module.vpc
         - mode: managed
         - type: aws_egress_only_internet_gateway
         - name: this
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 68
        
         - module: module.vpc
         - mode: managed
         - type: aws_eip
         - name: nat
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 0
                 - attributes
                    
                     - allocation_id: null
                     - associate_with_private_ip: null
                     - association_id: eipassoc-273d48db
                     - customer_owned_ip: 
                     - customer_owned_ipv4_pool: 
                     - domain: vpc
                     - id: eipalloc-08630ec82bcc9c83e
                     - instance: 
                     - network_interface: eni-020a408b2d5df3526
                     - private_dns: ip-10-0-1-96.us-west-1.compute.internal
                     - private_ip: 10.0.1.96
                     - public_dns: ec2-13-52-210-124.us-west-1.compute.amazonaws.com
                     - public_ip: 13.52.210.124
                     - public_ipv4_pool: amazon
                     - tags
                        
                         - Name: ibft4-us-west-1a
                         - terraform: true
                         - vpc: ibft4
                        
                     - timeouts: null
                     - vpc: true
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiZGVsZXRlIjoxODAwMDAwMDAwMDAsInJlYWQiOjkwMDAwMDAwMDAwMCwidXBkYXRlIjozMDAwMDAwMDAwMDB9fQ==
                
            
        
     - 69
        
         - module: module.vpc
         - mode: managed
         - type: aws_elasticache_subnet_group
         - name: elasticache
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 70
        
         - module: module.vpc
         - mode: managed
         - type: aws_flow_log
         - name: this
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 71
        
         - mode: managed
         - type: aws_iam_instance_profile
         - name: monitoring_profile
         - provider: provider.aws
         - instances
            
             - 0
                
                 - schema_version: 0
                 - attributes
                    
                     - create_date: 2020-06-14T14:31:01Z
                     - id: ibft4_us-west-1_monitoring_profile
                     - name: ibft4_us-west-1_monitoring_profile
                     - name_prefix: null
                     - path: /
                     - role: ibft4_us-west-1_monitoring_access_role
                     - roles
                        
                         - 0: ibft4_us-west-1_monitoring_access_role
                        
                     - unique_id: AIPAR6INBBYNQ2KP5QALT
                    
                 - private: bnVsbA==
                 - dependencies
                    
                     - 0: aws_iam_role.monitoring_access_role
                    
                
            
        
     - 72
        
         - module: module.vpc
         - mode: managed
         - type: aws_iam_policy
         - name: vpc_flow_log_cloudwatch
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 73
        
         - mode: managed
         - type: aws_iam_policy_attachment
         - name: monitoring_role_policy_attachment
         - provider: provider.aws
         - instances
            
             - 0
                
                 - schema_version: 0
                 - attributes
                    
                     - groups
                        
                        
                     - id: ibft4_us-west-1_monitoring_role_policy_attachment
                     - name: ibft4_us-west-1_monitoring_role_policy_attachment
                     - roles
                        
                         - 0: ibft4_us-west-1_monitoring_access_role
                        
                     - users
                        
                        
                    
                 - private: bnVsbA==
                 - dependencies
                    
                     - 0: aws_iam_role.monitoring_access_role
                    
                
            
        
     - 74
        
         - mode: managed
         - type: aws_iam_role
         - name: monitoring_access_role
         - provider: provider.aws
         - instances
            
             - 0
                
                 - schema_version: 0
                 - attributes
                    
                     - assume_role_policy: {"Version":"2012-10-17","Statement":[{"Sid":"","Effect":"Allow","Principal":{"Service":"ec2.amazonaws.com"},"Action":"sts:AssumeRole"}]}
                     - create_date: 2020-06-14T14:31:01Z
                     - description: 
                     - force_detach_policies: false
                     - id: ibft4_us-west-1_monitoring_access_role
                     - max_session_duration: 3600
                     - name: ibft4_us-west-1_monitoring_access_role
                     - name_prefix: null
                     - path: /
                     - permissions_boundary: null
                     - tags
                        
                        
                     - unique_id: AROAR6INBBYN54IUHLR6I
                    
                 - private: bnVsbA==
                
            
        
     - 75
        
         - module: module.vpc
         - mode: managed
         - type: aws_iam_role
         - name: vpc_flow_log_cloudwatch
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 76
        
         - module: module.vpc
         - mode: managed
         - type: aws_iam_role_policy_attachment
         - name: vpc_flow_log_cloudwatch
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 77
        
         - mode: managed
         - type: aws_instance
         - name: ibft_bootnode
         - provider: provider.aws
         - instances
            
             - 0
                
                 - schema_version: 1
                 - attributes
                    
                     - ami: ami-0d705db840ec5f0c5
                     - associate_public_ip_address: true
                     - availability_zone: us-west-1a
                     - cpu_core_count: 1
                     - cpu_threads_per_core: 2
                     - credit_specification
                        
                         - 0
                            
                             - cpu_credits: unlimited
                            
                        
                     - disable_api_termination: false
                     - ebs_block_device
                        
                        
                     - ebs_optimized: true
                     - ephemeral_block_device
                        
                        
                     - get_password_data: false
                     - hibernation: false
                     - host_id: null
                     - iam_instance_profile: ibft4_us-west-1_monitoring_profile
                     - id: i-087aed9a401bac9f1
                     - instance_initiated_shutdown_behavior: null
                     - instance_state: running
                     - instance_type: t3.medium
                     - ipv6_address_count: 0
                     - ipv6_addresses
                        
                        
                     - key_name: freight
                     - metadata_options
                        
                         - 0
                            
                             - http_endpoint: enabled
                             - http_put_response_hop_limit: 1
                             - http_tokens: optional
                            
                        
                     - monitoring: false
                     - network_interface
                        
                        
                     - network_interface_id: null
                     - password_data: 
                     - placement_group: 
                     - primary_network_interface_id: eni-00cbf679ab2112798
                     - private_dns: ip-10-0-1-9.us-west-1.compute.internal
                     - private_ip: 10.0.1.9
                     - public_dns: ec2-3-101-83-95.us-west-1.compute.amazonaws.com
                     - public_ip: 3.101.83.95
                     - root_block_device
                        
                         - 0
                            
                             - delete_on_termination: true
                             - device_name: /dev/sda1
                             - encrypted: false
                             - iops: 120
                             - kms_key_id: 
                             - volume_id: vol-0e9598480c410c017
                             - volume_size: 40
                             - volume_type: gp2
                            
                        
                     - security_groups
                        
                        
                     - source_dest_check: true
                     - subnet_id: subnet-07df858a84da18c00
                     - tags
                        
                         - Name: besu-ibft4-bootnode
                        
                     - tenancy: default
                     - timeouts: null
                     - user_data: null
                     - user_data_base64: null
                     - volume_tags
                        
                        
                     - vpc_security_group_ids
                        
                         - 0: sg-00497d9364b1d927b
                         - 1: sg-0fd0fd5fe13b9edca
                        
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6MTIwMDAwMDAwMDAwMCwidXBkYXRlIjo2MDAwMDAwMDAwMDB9LCJzY2hlbWFfdmVyc2lvbiI6IjEifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_security_group.eth_sg
                     - 2: module.ssh_security_group.module.sg.aws_security_group.this
                     - 3: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 4: module.vpc.aws_subnet.public
                    
                
            
        
     - 78
        
         - mode: managed
         - type: aws_instance
         - name: ibft_nodes
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 1
                 - attributes
                    
                     - ami: ami-0d705db840ec5f0c5
                     - associate_public_ip_address: true
                     - availability_zone: us-west-1a
                     - cpu_core_count: 1
                     - cpu_threads_per_core: 2
                     - credit_specification
                        
                         - 0
                            
                             - cpu_credits: unlimited
                            
                        
                     - disable_api_termination: false
                     - ebs_block_device
                        
                        
                     - ebs_optimized: true
                     - ephemeral_block_device
                        
                        
                     - get_password_data: false
                     - hibernation: false
                     - host_id: null
                     - iam_instance_profile: ibft4_us-west-1_monitoring_profile
                     - id: i-0d42bc81651d86a4a
                     - instance_initiated_shutdown_behavior: null
                     - instance_state: running
                     - instance_type: t3.medium
                     - ipv6_address_count: 0
                     - ipv6_addresses
                        
                        
                     - key_name: freight
                     - metadata_options
                        
                         - 0
                            
                             - http_endpoint: enabled
                             - http_put_response_hop_limit: 1
                             - http_tokens: optional
                            
                        
                     - monitoring: false
                     - network_interface
                        
                        
                     - network_interface_id: null
                     - password_data: 
                     - placement_group: 
                     - primary_network_interface_id: eni-02c61276a4f4c3c0e
                     - private_dns: ip-10-0-1-171.us-west-1.compute.internal
                     - private_ip: 10.0.1.171
                     - public_dns: ec2-54-183-133-228.us-west-1.compute.amazonaws.com
                     - public_ip: 54.183.133.228
                     - root_block_device
                        
                         - 0
                            
                             - delete_on_termination: true
                             - device_name: /dev/sda1
                             - encrypted: false
                             - iops: 120
                             - kms_key_id: 
                             - volume_id: vol-029b21fa8c72489ef
                             - volume_size: 40
                             - volume_type: gp2
                            
                        
                     - security_groups
                        
                        
                     - source_dest_check: true
                     - subnet_id: subnet-07df858a84da18c00
                     - tags
                        
                         - Name: besu-ibft4-0
                        
                     - tenancy: default
                     - timeouts: null
                     - user_data: null
                     - user_data_base64: null
                     - volume_tags
                        
                        
                     - vpc_security_group_ids
                        
                         - 0: sg-00497d9364b1d927b
                         - 1: sg-0fd0fd5fe13b9edca
                        
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6MTIwMDAwMDAwMDAwMCwidXBkYXRlIjo2MDAwMDAwMDAwMDB9LCJzY2hlbWFfdmVyc2lvbiI6IjEifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_instance.ibft_bootnode
                     - 2: aws_security_group.eth_sg
                     - 3: module.ssh_security_group.module.sg.aws_security_group.this
                     - 4: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 5: module.vpc.aws_subnet.public
                    
                
             - 1
                
                 - index_key: 1
                 - schema_version: 1
                 - attributes
                    
                     - ami: ami-0d705db840ec5f0c5
                     - associate_public_ip_address: true
                     - availability_zone: us-west-1c
                     - cpu_core_count: 1
                     - cpu_threads_per_core: 2
                     - credit_specification
                        
                         - 0
                            
                             - cpu_credits: unlimited
                            
                        
                     - disable_api_termination: false
                     - ebs_block_device
                        
                        
                     - ebs_optimized: true
                     - ephemeral_block_device
                        
                        
                     - get_password_data: false
                     - hibernation: false
                     - host_id: null
                     - iam_instance_profile: ibft4_us-west-1_monitoring_profile
                     - id: i-0d9c354e3c859e348
                     - instance_initiated_shutdown_behavior: null
                     - instance_state: running
                     - instance_type: t3.medium
                     - ipv6_address_count: 0
                     - ipv6_addresses
                        
                        
                     - key_name: freight
                     - metadata_options
                        
                         - 0
                            
                             - http_endpoint: enabled
                             - http_put_response_hop_limit: 1
                             - http_tokens: optional
                            
                        
                     - monitoring: false
                     - network_interface
                        
                        
                     - network_interface_id: null
                     - password_data: 
                     - placement_group: 
                     - primary_network_interface_id: eni-06f3373af80b9bf13
                     - private_dns: ip-10-0-2-89.us-west-1.compute.internal
                     - private_ip: 10.0.2.89
                     - public_dns: ec2-54-67-10-205.us-west-1.compute.amazonaws.com
                     - public_ip: 54.67.10.205
                     - root_block_device
                        
                         - 0
                            
                             - delete_on_termination: true
                             - device_name: /dev/sda1
                             - encrypted: false
                             - iops: 120
                             - kms_key_id: 
                             - volume_id: vol-0776a7915b5644894
                             - volume_size: 40
                             - volume_type: gp2
                            
                        
                     - security_groups
                        
                        
                     - source_dest_check: true
                     - subnet_id: subnet-0d619244f43aeea5e
                     - tags
                        
                         - Name: besu-ibft4-1
                        
                     - tenancy: default
                     - timeouts: null
                     - user_data: null
                     - user_data_base64: null
                     - volume_tags
                        
                        
                     - vpc_security_group_ids
                        
                         - 0: sg-00497d9364b1d927b
                         - 1: sg-0fd0fd5fe13b9edca
                        
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6MTIwMDAwMDAwMDAwMCwidXBkYXRlIjo2MDAwMDAwMDAwMDB9LCJzY2hlbWFfdmVyc2lvbiI6IjEifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_instance.ibft_bootnode
                     - 2: aws_security_group.eth_sg
                     - 3: module.ssh_security_group.module.sg.aws_security_group.this
                     - 4: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 5: module.vpc.aws_subnet.public
                    
                
             - 2
                
                 - index_key: 2
                 - schema_version: 1
                 - attributes
                    
                     - ami: ami-0d705db840ec5f0c5
                     - associate_public_ip_address: true
                     - availability_zone: us-west-1a
                     - cpu_core_count: 1
                     - cpu_threads_per_core: 2
                     - credit_specification
                        
                         - 0
                            
                             - cpu_credits: unlimited
                            
                        
                     - disable_api_termination: false
                     - ebs_block_device
                        
                        
                     - ebs_optimized: true
                     - ephemeral_block_device
                        
                        
                     - get_password_data: false
                     - hibernation: false
                     - host_id: null
                     - iam_instance_profile: ibft4_us-west-1_monitoring_profile
                     - id: i-01325fd7979104b4a
                     - instance_initiated_shutdown_behavior: null
                     - instance_state: running
                     - instance_type: t3.medium
                     - ipv6_address_count: 0
                     - ipv6_addresses
                        
                        
                     - key_name: freight
                     - metadata_options
                        
                         - 0
                            
                             - http_endpoint: enabled
                             - http_put_response_hop_limit: 1
                             - http_tokens: optional
                            
                        
                     - monitoring: false
                     - network_interface
                        
                        
                     - network_interface_id: null
                     - password_data: 
                     - placement_group: 
                     - primary_network_interface_id: eni-0b9ea74a922484f1e
                     - private_dns: ip-10-0-3-187.us-west-1.compute.internal
                     - private_ip: 10.0.3.187
                     - public_dns: ec2-54-193-83-95.us-west-1.compute.amazonaws.com
                     - public_ip: 54.193.83.95
                     - root_block_device
                        
                         - 0
                            
                             - delete_on_termination: true
                             - device_name: /dev/sda1
                             - encrypted: false
                             - iops: 120
                             - kms_key_id: 
                             - volume_id: vol-0e897bb21f0f2b822
                             - volume_size: 40
                             - volume_type: gp2
                            
                        
                     - security_groups
                        
                        
                     - source_dest_check: true
                     - subnet_id: subnet-0005aeeaff667e9f9
                     - tags
                        
                         - Name: besu-ibft4-2
                        
                     - tenancy: default
                     - timeouts: null
                     - user_data: null
                     - user_data_base64: null
                     - volume_tags
                        
                        
                     - vpc_security_group_ids
                        
                         - 0: sg-00497d9364b1d927b
                         - 1: sg-0fd0fd5fe13b9edca
                        
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6MTIwMDAwMDAwMDAwMCwidXBkYXRlIjo2MDAwMDAwMDAwMDB9LCJzY2hlbWFfdmVyc2lvbiI6IjEifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_instance.ibft_bootnode
                     - 2: aws_security_group.eth_sg
                     - 3: module.ssh_security_group.module.sg.aws_security_group.this
                     - 4: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 5: module.vpc.aws_subnet.public
                    
                
             - 3
                
                 - index_key: 3
                 - schema_version: 1
                 - attributes
                    
                     - ami: ami-0d705db840ec5f0c5
                     - associate_public_ip_address: true
                     - availability_zone: us-west-1a
                     - cpu_core_count: 1
                     - cpu_threads_per_core: 2
                     - credit_specification
                        
                         - 0
                            
                             - cpu_credits: unlimited
                            
                        
                     - disable_api_termination: false
                     - ebs_block_device
                        
                        
                     - ebs_optimized: true
                     - ephemeral_block_device
                        
                        
                     - get_password_data: false
                     - hibernation: false
                     - host_id: null
                     - iam_instance_profile: ibft4_us-west-1_monitoring_profile
                     - id: i-067ebe7878ad25122
                     - instance_initiated_shutdown_behavior: null
                     - instance_state: running
                     - instance_type: t3.medium
                     - ipv6_address_count: 0
                     - ipv6_addresses
                        
                        
                     - key_name: freight
                     - metadata_options
                        
                         - 0
                            
                             - http_endpoint: enabled
                             - http_put_response_hop_limit: 1
                             - http_tokens: optional
                            
                        
                     - monitoring: false
                     - network_interface
                        
                        
                     - network_interface_id: null
                     - password_data: 
                     - placement_group: 
                     - primary_network_interface_id: eni-04ee703fcb361b280
                     - private_dns: ip-10-0-1-170.us-west-1.compute.internal
                     - private_ip: 10.0.1.170
                     - public_dns: ec2-54-193-114-99.us-west-1.compute.amazonaws.com
                     - public_ip: 54.193.114.99
                     - root_block_device
                        
                         - 0
                            
                             - delete_on_termination: true
                             - device_name: /dev/sda1
                             - encrypted: false
                             - iops: 120
                             - kms_key_id: 
                             - volume_id: vol-0e4b41fbe292044ad
                             - volume_size: 40
                             - volume_type: gp2
                            
                        
                     - security_groups
                        
                        
                     - source_dest_check: true
                     - subnet_id: subnet-07df858a84da18c00
                     - tags
                        
                         - Name: besu-ibft4-3
                        
                     - tenancy: default
                     - timeouts: null
                     - user_data: null
                     - user_data_base64: null
                     - volume_tags
                        
                        
                     - vpc_security_group_ids
                        
                         - 0: sg-00497d9364b1d927b
                         - 1: sg-0fd0fd5fe13b9edca
                        
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6MTIwMDAwMDAwMDAwMCwidXBkYXRlIjo2MDAwMDAwMDAwMDB9LCJzY2hlbWFfdmVyc2lvbiI6IjEifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_instance.ibft_bootnode
                     - 2: aws_security_group.eth_sg
                     - 3: module.ssh_security_group.module.sg.aws_security_group.this
                     - 4: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 5: module.vpc.aws_subnet.public
                    
                
             - 4
                
                 - index_key: 4
                 - schema_version: 1
                 - attributes
                    
                     - ami: ami-0d705db840ec5f0c5
                     - associate_public_ip_address: true
                     - availability_zone: us-west-1c
                     - cpu_core_count: 1
                     - cpu_threads_per_core: 2
                     - credit_specification
                        
                         - 0
                            
                             - cpu_credits: unlimited
                            
                        
                     - disable_api_termination: false
                     - ebs_block_device
                        
                        
                     - ebs_optimized: true
                     - ephemeral_block_device
                        
                        
                     - get_password_data: false
                     - hibernation: false
                     - host_id: null
                     - iam_instance_profile: ibft4_us-west-1_monitoring_profile
                     - id: i-003b99d6d7f6aef63
                     - instance_initiated_shutdown_behavior: null
                     - instance_state: running
                     - instance_type: t3.medium
                     - ipv6_address_count: 0
                     - ipv6_addresses
                        
                        
                     - key_name: freight
                     - metadata_options
                        
                         - 0
                            
                             - http_endpoint: enabled
                             - http_put_response_hop_limit: 1
                             - http_tokens: optional
                            
                        
                     - monitoring: false
                     - network_interface
                        
                        
                     - network_interface_id: null
                     - password_data: 
                     - placement_group: 
                     - primary_network_interface_id: eni-08f3949fd2db19895
                     - private_dns: ip-10-0-2-228.us-west-1.compute.internal
                     - private_ip: 10.0.2.228
                     - public_dns: ec2-54-183-62-83.us-west-1.compute.amazonaws.com
                     - public_ip: 54.183.62.83
                     - root_block_device
                        
                         - 0
                            
                             - delete_on_termination: true
                             - device_name: /dev/sda1
                             - encrypted: false
                             - iops: 120
                             - kms_key_id: 
                             - volume_id: vol-0714c0dea726979bc
                             - volume_size: 40
                             - volume_type: gp2
                            
                        
                     - security_groups
                        
                        
                     - source_dest_check: true
                     - subnet_id: subnet-0d619244f43aeea5e
                     - tags
                        
                         - Name: besu-ibft4-4
                        
                     - tenancy: default
                     - timeouts: null
                     - user_data: null
                     - user_data_base64: null
                     - volume_tags
                        
                        
                     - vpc_security_group_ids
                        
                         - 0: sg-00497d9364b1d927b
                         - 1: sg-0fd0fd5fe13b9edca
                        
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6MTIwMDAwMDAwMDAwMCwidXBkYXRlIjo2MDAwMDAwMDAwMDB9LCJzY2hlbWFfdmVyc2lvbiI6IjEifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_instance.ibft_bootnode
                     - 2: aws_security_group.eth_sg
                     - 3: module.ssh_security_group.module.sg.aws_security_group.this
                     - 4: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 5: module.vpc.aws_subnet.public
                    
                
            
        
     - 79
        
         - mode: managed
         - type: aws_instance
         - name: ibft_rpcnode
         - provider: provider.aws
         - instances
            
             - 0
                
                 - schema_version: 1
                 - attributes
                    
                     - ami: ami-0d705db840ec5f0c5
                     - associate_public_ip_address: true
                     - availability_zone: us-west-1c
                     - cpu_core_count: 1
                     - cpu_threads_per_core: 2
                     - credit_specification
                        
                         - 0
                            
                             - cpu_credits: unlimited
                            
                        
                     - disable_api_termination: false
                     - ebs_block_device
                        
                        
                     - ebs_optimized: true
                     - ephemeral_block_device
                        
                        
                     - get_password_data: false
                     - hibernation: false
                     - host_id: null
                     - iam_instance_profile: ibft4_us-west-1_monitoring_profile
                     - id: i-081b46211db44b018
                     - instance_initiated_shutdown_behavior: null
                     - instance_state: running
                     - instance_type: t3.medium
                     - ipv6_address_count: 0
                     - ipv6_addresses
                        
                        
                     - key_name: freight
                     - metadata_options
                        
                         - 0
                            
                             - http_endpoint: enabled
                             - http_put_response_hop_limit: 1
                             - http_tokens: optional
                            
                        
                     - monitoring: false
                     - network_interface
                        
                        
                     - network_interface_id: null
                     - password_data: 
                     - placement_group: 
                     - primary_network_interface_id: eni-021c878ed4410115f
                     - private_dns: ip-10-0-2-71.us-west-1.compute.internal
                     - private_ip: 10.0.2.71
                     - public_dns: ec2-54-193-216-230.us-west-1.compute.amazonaws.com
                     - public_ip: 54.193.216.230
                     - root_block_device
                        
                         - 0
                            
                             - delete_on_termination: true
                             - device_name: /dev/sda1
                             - encrypted: false
                             - iops: 120
                             - kms_key_id: 
                             - volume_id: vol-0a560306cd188d25e
                             - volume_size: 40
                             - volume_type: gp2
                            
                        
                     - security_groups
                        
                        
                     - source_dest_check: true
                     - subnet_id: subnet-0d619244f43aeea5e
                     - tags
                        
                         - Name: besu-ibft4-rpcnode
                        
                     - tenancy: default
                     - timeouts: null
                     - user_data: null
                     - user_data_base64: null
                     - volume_tags
                        
                        
                     - vpc_security_group_ids
                        
                         - 0: sg-00497d9364b1d927b
                         - 1: sg-0fd0fd5fe13b9edca
                        
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6MTIwMDAwMDAwMDAwMCwidXBkYXRlIjo2MDAwMDAwMDAwMDB9LCJzY2hlbWFfdmVyc2lvbiI6IjEifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_instance.ibft_bootnode
                     - 2: aws_security_group.eth_sg
                     - 3: module.ssh_security_group.module.sg.aws_security_group.this
                     - 4: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 5: module.vpc.aws_subnet.public
                    
                
            
        
     - 80
        
         - mode: managed
         - type: aws_instance
         - name: monitoring
         - provider: provider.aws
         - instances
            
             - 0
                
                 - schema_version: 1
                 - attributes
                    
                     - ami: ami-0d705db840ec5f0c5
                     - associate_public_ip_address: true
                     - availability_zone: us-west-1a
                     - cpu_core_count: 1
                     - cpu_threads_per_core: 2
                     - credit_specification
                        
                         - 0
                            
                             - cpu_credits: unlimited
                            
                        
                     - disable_api_termination: false
                     - ebs_block_device
                        
                        
                     - ebs_optimized: true
                     - ephemeral_block_device
                        
                        
                     - get_password_data: false
                     - hibernation: false
                     - host_id: null
                     - iam_instance_profile: ibft4_us-west-1_monitoring_profile
                     - id: i-0c1d2c2c11fa0667e
                     - instance_initiated_shutdown_behavior: null
                     - instance_state: running
                     - instance_type: t3.micro
                     - ipv6_address_count: 0
                     - ipv6_addresses
                        
                        
                     - key_name: freight
                     - metadata_options
                        
                         - 0
                            
                             - http_endpoint: enabled
                             - http_put_response_hop_limit: 1
                             - http_tokens: optional
                            
                        
                     - monitoring: false
                     - network_interface
                        
                        
                     - network_interface_id: null
                     - password_data: 
                     - placement_group: 
                     - primary_network_interface_id: eni-044910e3a13d74471
                     - private_dns: ip-10-0-1-64.us-west-1.compute.internal
                     - private_ip: 10.0.1.64
                     - public_dns: ec2-18-144-58-233.us-west-1.compute.amazonaws.com
                     - public_ip: 18.144.58.233
                     - root_block_device
                        
                         - 0
                            
                             - delete_on_termination: true
                             - device_name: /dev/sda1
                             - encrypted: false
                             - iops: 120
                             - kms_key_id: 
                             - volume_id: vol-0a267fe1ce900b328
                             - volume_size: 40
                             - volume_type: gp2
                            
                        
                     - security_groups
                        
                        
                     - source_dest_check: true
                     - subnet_id: subnet-07df858a84da18c00
                     - tags
                        
                         - Name: ibft4-monitoring
                        
                     - tenancy: default
                     - timeouts: null
                     - user_data: null
                     - user_data_base64: null
                     - volume_tags
                        
                        
                     - vpc_security_group_ids
                        
                         - 0: sg-0a5949569274ac766
                         - 1: sg-0fd0fd5fe13b9edca
                        
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6MTIwMDAwMDAwMDAwMCwidXBkYXRlIjo2MDAwMDAwMDAwMDB9LCJzY2hlbWFfdmVyc2lvbiI6IjEifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_security_group.monitoring_sg
                     - 2: module.ssh_security_group.module.sg.aws_security_group.this
                     - 3: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 4: module.vpc.aws_subnet.public
                    
                
            
        
     - 81
        
         - module: module.vpc
         - mode: managed
         - type: aws_internet_gateway
         - name: this
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 0
                 - attributes
                    
                     - id: igw-09f07d54b09cfa46e
                     - owner_id: 133708189211
                     - tags
                        
                         - Name: ibft4
                         - terraform: true
                         - vpc: ibft4
                        
                     - vpc_id: vpc-0881077c57fc53714
                    
                 - private: bnVsbA==
                 - dependencies
                    
                     - 0: module.vpc.aws_vpc.this
                     - 1: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
            
        
     - 82
        
         - module: module.vpc
         - mode: managed
         - type: aws_nat_gateway
         - name: this
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 0
                 - attributes
                    
                     - allocation_id: eipalloc-08630ec82bcc9c83e
                     - id: nat-0307203d58bc4dd0a
                     - network_interface_id: eni-020a408b2d5df3526
                     - private_ip: 10.0.1.96
                     - public_ip: 13.52.210.124
                     - subnet_id: subnet-07df858a84da18c00
                     - tags
                        
                         - Name: ibft4-us-west-1a
                         - terraform: true
                         - vpc: ibft4
                        
                    
                 - private: bnVsbA==
                 - dependencies
                    
                     - 0: module.vpc.aws_eip.nat
                     - 1: module.vpc.aws_internet_gateway.this
                     - 2: module.vpc.aws_subnet.public
                     - 3: module.vpc.aws_vpc.this
                     - 4: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
            
        
     - 83
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl
         - name: database
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 84
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl
         - name: elasticache
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 85
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl
         - name: intra
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 86
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl
         - name: private
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 87
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl
         - name: public
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 88
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl
         - name: redshift
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 89
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl_rule
         - name: database_inbound
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 90
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl_rule
         - name: database_outbound
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 91
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl_rule
         - name: elasticache_inbound
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 92
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl_rule
         - name: elasticache_outbound
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 93
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl_rule
         - name: intra_inbound
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 94
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl_rule
         - name: intra_outbound
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 95
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl_rule
         - name: private_inbound
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 96
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl_rule
         - name: private_outbound
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 97
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl_rule
         - name: public_inbound
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 98
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl_rule
         - name: public_outbound
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 99
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl_rule
         - name: redshift_inbound
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 100
        
         - module: module.vpc
         - mode: managed
         - type: aws_network_acl_rule
         - name: redshift_outbound
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 101
        
         - module: module.vpc
         - mode: managed
         - type: aws_redshift_subnet_group
         - name: redshift
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 102
        
         - module: module.vpc
         - mode: managed
         - type: aws_route
         - name: database_internet_gateway
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 103
        
         - module: module.vpc
         - mode: managed
         - type: aws_route
         - name: database_ipv6_egress
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 104
        
         - module: module.vpc
         - mode: managed
         - type: aws_route
         - name: database_nat_gateway
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 105
        
         - module: module.vpc
         - mode: managed
         - type: aws_route
         - name: private_ipv6_egress
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 106
        
         - module: module.vpc
         - mode: managed
         - type: aws_route
         - name: private_nat_gateway
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 0
                 - attributes
                    
                     - destination_cidr_block: 0.0.0.0/0
                     - destination_ipv6_cidr_block: null
                     - destination_prefix_list_id: 
                     - egress_only_gateway_id: 
                     - gateway_id: 
                     - id: r-rtb-09eb4486e863c6ac91080289494
                     - instance_id: 
                     - instance_owner_id: 
                     - nat_gateway_id: nat-0307203d58bc4dd0a
                     - network_interface_id: 
                     - origin: CreateRoute
                     - route_table_id: rtb-09eb4486e863c6ac9
                     - state: active
                     - timeouts
                        
                         - create: 5m
                         - delete: null
                        
                     - transit_gateway_id: 
                     - vpc_peering_connection_id: 
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjozMDAwMDAwMDAwMDAsImRlbGV0ZSI6MzAwMDAwMDAwMDAwfX0=
                 - dependencies
                    
                     - 0: module.vpc.aws_eip.nat
                     - 1: module.vpc.aws_internet_gateway.this
                     - 2: module.vpc.aws_nat_gateway.this
                     - 3: module.vpc.aws_route_table.private
                     - 4: module.vpc.aws_subnet.public
                     - 5: module.vpc.aws_vpc.this
                     - 6: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
            
        
     - 107
        
         - module: module.vpc
         - mode: managed
         - type: aws_route
         - name: public_internet_gateway
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 0
                 - attributes
                    
                     - destination_cidr_block: 0.0.0.0/0
                     - destination_ipv6_cidr_block: null
                     - destination_prefix_list_id: 
                     - egress_only_gateway_id: 
                     - gateway_id: igw-09f07d54b09cfa46e
                     - id: r-rtb-08a9f3ea4ceb9c7391080289494
                     - instance_id: 
                     - instance_owner_id: 
                     - nat_gateway_id: 
                     - network_interface_id: 
                     - origin: CreateRoute
                     - route_table_id: rtb-08a9f3ea4ceb9c739
                     - state: active
                     - timeouts
                        
                         - create: 5m
                         - delete: null
                        
                     - transit_gateway_id: 
                     - vpc_peering_connection_id: 
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjozMDAwMDAwMDAwMDAsImRlbGV0ZSI6MzAwMDAwMDAwMDAwfX0=
                 - dependencies
                    
                     - 0: module.vpc.aws_internet_gateway.this
                     - 1: module.vpc.aws_route_table.public
                     - 2: module.vpc.aws_vpc.this
                     - 3: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
            
        
     - 108
        
         - module: module.vpc
         - mode: managed
         - type: aws_route
         - name: public_internet_gateway_ipv6
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 109
        
         - mode: managed
         - type: aws_route53_record
         - name: bootnode_dns
         - provider: provider.aws
         - instances
            
             - 0
                
                 - schema_version: 2
                 - attributes
                    
                     - alias
                        
                        
                     - allow_overwrite: null
                     - failover_routing_policy
                        
                        
                     - fqdn: bootnode.ibft4.us-west-1
                     - geolocation_routing_policy
                        
                        
                     - health_check_id: null
                     - id: Z01635402A14L7Q8IKWW0_bootnode.ibft4.us-west-1_A
                     - latency_routing_policy
                        
                        
                     - multivalue_answer_routing_policy: null
                     - name: bootnode.ibft4.us-west-1
                     - records
                        
                         - 0: 10.0.1.9
                        
                     - set_identifier: null
                     - ttl: 300
                     - type: A
                     - weighted_routing_policy
                        
                        
                     - zone_id: Z01635402A14L7Q8IKWW0
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_instance.ibft_bootnode
                     - 2: aws_route53_zone.private
                     - 3: aws_security_group.eth_sg
                     - 4: module.ssh_security_group.module.sg.aws_security_group.this
                     - 5: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 6: module.vpc.aws_subnet.public
                    
                
            
        
     - 110
        
         - mode: managed
         - type: aws_route53_record
         - name: monitoring_dns
         - provider: provider.aws
         - instances
            
             - 0
                
                 - schema_version: 2
                 - attributes
                    
                     - alias
                        
                        
                     - allow_overwrite: null
                     - failover_routing_policy
                        
                        
                     - fqdn: monitoring.ibft4.us-west-1
                     - geolocation_routing_policy
                        
                        
                     - health_check_id: null
                     - id: Z01635402A14L7Q8IKWW0_monitoring.ibft4.us-west-1_A
                     - latency_routing_policy
                        
                        
                     - multivalue_answer_routing_policy: null
                     - name: monitoring.ibft4.us-west-1
                     - records
                        
                         - 0: 10.0.1.64
                        
                     - set_identifier: null
                     - ttl: 300
                     - type: A
                     - weighted_routing_policy
                        
                        
                     - zone_id: Z01635402A14L7Q8IKWW0
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_instance.monitoring
                     - 2: aws_route53_zone.private
                     - 3: aws_security_group.monitoring_sg
                     - 4: module.ssh_security_group.module.sg.aws_security_group.this
                     - 5: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 6: module.vpc.aws_subnet.public
                    
                
            
        
     - 111
        
         - mode: managed
         - type: aws_route53_record
         - name: nodes_dns
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 2
                 - attributes
                    
                     - alias
                        
                        
                     - allow_overwrite: null
                     - failover_routing_policy
                        
                        
                     - fqdn: node-0.ibft4.us-west-1
                     - geolocation_routing_policy
                        
                        
                     - health_check_id: null
                     - id: Z01635402A14L7Q8IKWW0_node-0.ibft4.us-west-1_A
                     - latency_routing_policy
                        
                        
                     - multivalue_answer_routing_policy: null
                     - name: node-0.ibft4.us-west-1
                     - records
                        
                         - 0: 10.0.1.171
                        
                     - set_identifier: null
                     - ttl: 300
                     - type: A
                     - weighted_routing_policy
                        
                        
                     - zone_id: Z01635402A14L7Q8IKWW0
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_instance.ibft_bootnode
                     - 2: aws_instance.ibft_nodes
                     - 3: aws_route53_zone.private
                     - 4: aws_security_group.eth_sg
                     - 5: module.ssh_security_group.module.sg.aws_security_group.this
                     - 6: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 7: module.vpc.aws_subnet.public
                    
                
             - 1
                
                 - index_key: 1
                 - schema_version: 2
                 - attributes
                    
                     - alias
                        
                        
                     - allow_overwrite: null
                     - failover_routing_policy
                        
                        
                     - fqdn: node-1.ibft4.us-west-1
                     - geolocation_routing_policy
                        
                        
                     - health_check_id: null
                     - id: Z01635402A14L7Q8IKWW0_node-1.ibft4.us-west-1_A
                     - latency_routing_policy
                        
                        
                     - multivalue_answer_routing_policy: null
                     - name: node-1.ibft4.us-west-1
                     - records
                        
                         - 0: 10.0.2.89
                        
                     - set_identifier: null
                     - ttl: 300
                     - type: A
                     - weighted_routing_policy
                        
                        
                     - zone_id: Z01635402A14L7Q8IKWW0
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_instance.ibft_bootnode
                     - 2: aws_instance.ibft_nodes
                     - 3: aws_route53_zone.private
                     - 4: aws_security_group.eth_sg
                     - 5: module.ssh_security_group.module.sg.aws_security_group.this
                     - 6: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 7: module.vpc.aws_subnet.public
                    
                
             - 2
                
                 - index_key: 2
                 - schema_version: 2
                 - attributes
                    
                     - alias
                        
                        
                     - allow_overwrite: null
                     - failover_routing_policy
                        
                        
                     - fqdn: node-2.ibft4.us-west-1
                     - geolocation_routing_policy
                        
                        
                     - health_check_id: null
                     - id: Z01635402A14L7Q8IKWW0_node-2.ibft4.us-west-1_A
                     - latency_routing_policy
                        
                        
                     - multivalue_answer_routing_policy: null
                     - name: node-2.ibft4.us-west-1
                     - records
                        
                         - 0: 10.0.3.187
                        
                     - set_identifier: null
                     - ttl: 300
                     - type: A
                     - weighted_routing_policy
                        
                        
                     - zone_id: Z01635402A14L7Q8IKWW0
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_instance.ibft_bootnode
                     - 2: aws_instance.ibft_nodes
                     - 3: aws_route53_zone.private
                     - 4: aws_security_group.eth_sg
                     - 5: module.ssh_security_group.module.sg.aws_security_group.this
                     - 6: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 7: module.vpc.aws_subnet.public
                    
                
             - 3
                
                 - index_key: 3
                 - schema_version: 2
                 - attributes
                    
                     - alias
                        
                        
                     - allow_overwrite: null
                     - failover_routing_policy
                        
                        
                     - fqdn: node-3.ibft4.us-west-1
                     - geolocation_routing_policy
                        
                        
                     - health_check_id: null
                     - id: Z01635402A14L7Q8IKWW0_node-3.ibft4.us-west-1_A
                     - latency_routing_policy
                        
                        
                     - multivalue_answer_routing_policy: null
                     - name: node-3.ibft4.us-west-1
                     - records
                        
                         - 0: 10.0.1.170
                        
                     - set_identifier: null
                     - ttl: 300
                     - type: A
                     - weighted_routing_policy
                        
                        
                     - zone_id: Z01635402A14L7Q8IKWW0
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_instance.ibft_bootnode
                     - 2: aws_instance.ibft_nodes
                     - 3: aws_route53_zone.private
                     - 4: aws_security_group.eth_sg
                     - 5: module.ssh_security_group.module.sg.aws_security_group.this
                     - 6: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 7: module.vpc.aws_subnet.public
                    
                
             - 4
                
                 - index_key: 4
                 - schema_version: 2
                 - attributes
                    
                     - alias
                        
                        
                     - allow_overwrite: null
                     - failover_routing_policy
                        
                        
                     - fqdn: node-4.ibft4.us-west-1
                     - geolocation_routing_policy
                        
                        
                     - health_check_id: null
                     - id: Z01635402A14L7Q8IKWW0_node-4.ibft4.us-west-1_A
                     - latency_routing_policy
                        
                        
                     - multivalue_answer_routing_policy: null
                     - name: node-4.ibft4.us-west-1
                     - records
                        
                         - 0: 10.0.2.228
                        
                     - set_identifier: null
                     - ttl: 300
                     - type: A
                     - weighted_routing_policy
                        
                        
                     - zone_id: Z01635402A14L7Q8IKWW0
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_instance.ibft_bootnode
                     - 2: aws_instance.ibft_nodes
                     - 3: aws_route53_zone.private
                     - 4: aws_security_group.eth_sg
                     - 5: module.ssh_security_group.module.sg.aws_security_group.this
                     - 6: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 7: module.vpc.aws_subnet.public
                    
                
            
        
     - 112
        
         - mode: managed
         - type: aws_route53_record
         - name: rpcnode_dns
         - provider: provider.aws
         - instances
            
             - 0
                
                 - schema_version: 2
                 - attributes
                    
                     - alias
                        
                        
                     - allow_overwrite: null
                     - failover_routing_policy
                        
                        
                     - fqdn: rpcnode.ibft4.us-west-1
                     - geolocation_routing_policy
                        
                        
                     - health_check_id: null
                     - id: Z01635402A14L7Q8IKWW0_rpcnode.ibft4.us-west-1_A
                     - latency_routing_policy
                        
                        
                     - multivalue_answer_routing_policy: null
                     - name: rpcnode.ibft4.us-west-1
                     - records
                        
                         - 0: 10.0.2.71
                        
                     - set_identifier: null
                     - ttl: 300
                     - type: A
                     - weighted_routing_policy
                        
                        
                     - zone_id: Z01635402A14L7Q8IKWW0
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: aws_iam_instance_profile.monitoring_profile
                     - 1: aws_instance.ibft_bootnode
                     - 2: aws_instance.ibft_rpcnode
                     - 3: aws_route53_zone.private
                     - 4: aws_security_group.eth_sg
                     - 5: module.ssh_security_group.module.sg.aws_security_group.this
                     - 6: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 7: module.vpc.aws_subnet.public
                    
                
            
        
     - 113
        
         - mode: managed
         - type: aws_route53_zone
         - name: private
         - provider: provider.aws
         - instances
            
             - 0
                
                 - schema_version: 0
                 - attributes
                    
                     - comment: Managed by Terraform
                     - delegation_set_id: 
                     - force_destroy: false
                     - id: Z01635402A14L7Q8IKWW0
                     - name: ibft4.us-west-1.
                     - name_servers
                        
                         - 0: ns-0.awsdns-00.com.
                         - 1: ns-1024.awsdns-00.org.
                         - 2: ns-1536.awsdns-00.co.uk.
                         - 3: ns-512.awsdns-00.net.
                        
                     - tags
                        
                        
                     - vpc
                        
                         - 0
                            
                             - vpc_id: vpc-0881077c57fc53714
                             - vpc_region: us-west-1
                            
                        
                     - vpc_id: null
                     - vpc_region: null
                     - zone_id: Z01635402A14L7Q8IKWW0
                    
                 - private: bnVsbA==
                 - dependencies
                    
                     - 0: module.vpc.aws_vpc.this
                    
                
            
        
     - 114
        
         - module: module.vpc
         - mode: managed
         - type: aws_route_table
         - name: database
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 115
        
         - module: module.vpc
         - mode: managed
         - type: aws_route_table
         - name: elasticache
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 116
        
         - module: module.vpc
         - mode: managed
         - type: aws_route_table
         - name: intra
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 117
        
         - module: module.vpc
         - mode: managed
         - type: aws_route_table
         - name: private
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 0
                 - attributes
                    
                     - id: rtb-09eb4486e863c6ac9
                     - owner_id: 133708189211
                     - propagating_vgws
                        
                        
                     - route
                        
                         - 0
                            
                             - cidr_block: 0.0.0.0/0
                             - egress_only_gateway_id: 
                             - gateway_id: 
                             - instance_id: 
                             - ipv6_cidr_block: 
                             - nat_gateway_id: nat-0307203d58bc4dd0a
                             - network_interface_id: 
                             - transit_gateway_id: 
                             - vpc_peering_connection_id: 
                            
                        
                     - tags
                        
                         - Name: ibft4-private
                         - terraform: true
                         - vpc: ibft4
                        
                     - vpc_id: vpc-0881077c57fc53714
                    
                 - private: bnVsbA==
                 - dependencies
                    
                     - 0: module.vpc.aws_vpc.this
                     - 1: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
            
        
     - 118
        
         - module: module.vpc
         - mode: managed
         - type: aws_route_table
         - name: public
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 0
                 - attributes
                    
                     - id: rtb-08a9f3ea4ceb9c739
                     - owner_id: 133708189211
                     - propagating_vgws
                        
                        
                     - route
                        
                         - 0
                            
                             - cidr_block: 0.0.0.0/0
                             - egress_only_gateway_id: 
                             - gateway_id: igw-09f07d54b09cfa46e
                             - instance_id: 
                             - ipv6_cidr_block: 
                             - nat_gateway_id: 
                             - network_interface_id: 
                             - transit_gateway_id: 
                             - vpc_peering_connection_id: 
                            
                        
                     - tags
                        
                         - Name: ibft4-public
                         - terraform: true
                         - vpc: ibft4
                        
                     - vpc_id: vpc-0881077c57fc53714
                    
                 - private: bnVsbA==
                 - dependencies
                    
                     - 0: module.vpc.aws_vpc.this
                     - 1: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
            
        
     - 119
        
         - module: module.vpc
         - mode: managed
         - type: aws_route_table
         - name: redshift
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 120
        
         - module: module.vpc
         - mode: managed
         - type: aws_route_table_association
         - name: database
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 121
        
         - module: module.vpc
         - mode: managed
         - type: aws_route_table_association
         - name: elasticache
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 122
        
         - module: module.vpc
         - mode: managed
         - type: aws_route_table_association
         - name: intra
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 123
        
         - module: module.vpc
         - mode: managed
         - type: aws_route_table_association
         - name: private
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 0
                 - attributes
                    
                     - gateway_id: null
                     - id: rtbassoc-06ff94fb7a93dbe05
                     - route_table_id: rtb-09eb4486e863c6ac9
                     - subnet_id: subnet-06e62b6d0c1d8e1b0
                    
                 - private: bnVsbA==
                 - dependencies
                    
                     - 0: module.vpc.aws_route_table.private
                     - 1: module.vpc.aws_subnet.private
                     - 2: module.vpc.aws_vpc.this
                     - 3: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
             - 1
                
                 - index_key: 1
                 - schema_version: 0
                 - attributes
                    
                     - gateway_id: null
                     - id: rtbassoc-0fb203e0a7444337b
                     - route_table_id: rtb-09eb4486e863c6ac9
                     - subnet_id: subnet-003fdf10e9015f50a
                    
                 - private: bnVsbA==
                 - dependencies
                    
                     - 0: module.vpc.aws_route_table.private
                     - 1: module.vpc.aws_subnet.private
                     - 2: module.vpc.aws_vpc.this
                     - 3: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
             - 2
                
                 - index_key: 2
                 - schema_version: 0
                 - attributes
                    
                     - gateway_id: null
                     - id: rtbassoc-07d9766998b5868c8
                     - route_table_id: rtb-09eb4486e863c6ac9
                     - subnet_id: subnet-0d97d5bad724ebcac
                    
                 - private: bnVsbA==
                 - dependencies
                    
                     - 0: module.vpc.aws_route_table.private
                     - 1: module.vpc.aws_subnet.private
                     - 2: module.vpc.aws_vpc.this
                     - 3: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
            
        
     - 124
        
         - module: module.vpc
         - mode: managed
         - type: aws_route_table_association
         - name: public
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 0
                 - attributes
                    
                     - gateway_id: null
                     - id: rtbassoc-07fc4836e4a988e97
                     - route_table_id: rtb-08a9f3ea4ceb9c739
                     - subnet_id: subnet-07df858a84da18c00
                    
                 - private: bnVsbA==
                 - dependencies
                    
                     - 0: module.vpc.aws_route_table.public
                     - 1: module.vpc.aws_subnet.public
                     - 2: module.vpc.aws_vpc.this
                     - 3: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
             - 1
                
                 - index_key: 1
                 - schema_version: 0
                 - attributes
                    
                     - gateway_id: null
                     - id: rtbassoc-01bfe4d6e80b89282
                     - route_table_id: rtb-08a9f3ea4ceb9c739
                     - subnet_id: subnet-0d619244f43aeea5e
                    
                 - private: bnVsbA==
                 - dependencies
                    
                     - 0: module.vpc.aws_route_table.public
                     - 1: module.vpc.aws_subnet.public
                     - 2: module.vpc.aws_vpc.this
                     - 3: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
             - 2
                
                 - index_key: 2
                 - schema_version: 0
                 - attributes
                    
                     - gateway_id: null
                     - id: rtbassoc-0e1bd94b7d5842f82
                     - route_table_id: rtb-08a9f3ea4ceb9c739
                     - subnet_id: subnet-0005aeeaff667e9f9
                    
                 - private: bnVsbA==
                 - dependencies
                    
                     - 0: module.vpc.aws_route_table.public
                     - 1: module.vpc.aws_subnet.public
                     - 2: module.vpc.aws_vpc.this
                     - 3: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
            
        
     - 125
        
         - module: module.vpc
         - mode: managed
         - type: aws_route_table_association
         - name: redshift
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 126
        
         - module: module.vpc
         - mode: managed
         - type: aws_route_table_association
         - name: redshift_public
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 127
        
         - mode: managed
         - type: aws_security_group
         - name: eth_sg
         - provider: provider.aws
         - instances
            
             - 0
                
                 - schema_version: 1
                 - attributes
                    
                     - description: ibft4_eth_sg
                     - egress
                        
                         - 0
                            
                             - cidr_blocks
                                
                                 - 0: 0.0.0.0/0
                                
                             - description: 
                             - from_port: 0
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: -1
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 0
                            
                        
                     - id: sg-00497d9364b1d927b
                     - ingress
                        
                         - 0
                            
                             - cidr_blocks
                                
                                 - 0: 10.0.0.0/16
                                
                             - description: 
                             - from_port: 30303
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: tcp
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 30303
                            
                         - 1
                            
                             - cidr_blocks
                                
                                 - 0: 10.0.0.0/16
                                
                             - description: 
                             - from_port: 30303
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: udp
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 30303
                            
                         - 2
                            
                             - cidr_blocks
                                
                                 - 0: 10.0.0.0/16
                                
                             - description: 
                             - from_port: 8545
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: tcp
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 8545
                            
                         - 3
                            
                             - cidr_blocks
                                
                                 - 0: 10.0.0.0/16
                                
                             - description: 
                             - from_port: 8546
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: tcp
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 8546
                            
                         - 4
                            
                             - cidr_blocks
                                
                                 - 0: 10.0.0.0/16
                                
                             - description: 
                             - from_port: 8547
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: tcp
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 8547
                            
                         - 5
                            
                             - cidr_blocks
                                
                                 - 0: 10.0.0.0/16
                                
                             - description: 
                             - from_port: 9100
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: tcp
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 9100
                            
                         - 6
                            
                             - cidr_blocks
                                
                                 - 0: 10.0.0.0/16
                                
                             - description: 
                             - from_port: 9545
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: tcp
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 9545
                            
                        
                     - name: ibft4_eth_sg
                     - name_prefix: null
                     - owner_id: 133708189211
                     - revoke_rules_on_delete: false
                     - tags
                        
                        
                     - timeouts: null
                     - vpc_id: vpc-0881077c57fc53714
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6NjAwMDAwMDAwMDAwfSwic2NoZW1hX3ZlcnNpb24iOiIxIn0=
                 - dependencies
                    
                     - 0: module.vpc.aws_vpc.this
                    
                
            
        
     - 128
        
         - mode: managed
         - type: aws_security_group
         - name: monitoring_sg
         - provider: provider.aws
         - instances
            
             - 0
                
                 - schema_version: 1
                 - attributes
                    
                     - description: ibft4_monitoring_sg
                     - egress
                        
                         - 0
                            
                             - cidr_blocks
                                
                                 - 0: 0.0.0.0/0
                                
                             - description: 
                             - from_port: 0
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: -1
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 0
                            
                        
                     - id: sg-0a5949569274ac766
                     - ingress
                        
                         - 0
                            
                             - cidr_blocks
                                
                                 - 0: 0.0.0.0/0
                                
                             - description: 
                             - from_port: 3000
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: tcp
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 3000
                            
                         - 1
                            
                             - cidr_blocks
                                
                                 - 0: 10.0.0.0/16
                                
                             - description: 
                             - from_port: 9090
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: tcp
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 9090
                            
                        
                     - name: ibft4_monitoring_sg
                     - name_prefix: null
                     - owner_id: 133708189211
                     - revoke_rules_on_delete: false
                     - tags
                        
                        
                     - timeouts: null
                     - vpc_id: vpc-0881077c57fc53714
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6NjAwMDAwMDAwMDAwfSwic2NoZW1hX3ZlcnNpb24iOiIxIn0=
                 - dependencies
                    
                     - 0: module.vpc.aws_vpc.this
                    
                
            
        
     - 129
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group
         - name: this
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 130
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group
         - name: this
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 131
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group
         - name: this
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 132
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group
         - name: this_name_prefix
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 1
                 - attributes
                    
                     - description: ibft4_http_sg
                     - egress
                        
                         - 0
                            
                             - cidr_blocks
                                
                                 - 0: 0.0.0.0/0
                                
                             - description: All protocols
                             - from_port: 0
                             - ipv6_cidr_blocks
                                
                                 - 0: ::/0
                                
                             - prefix_list_ids
                                
                                
                             - protocol: -1
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 0
                            
                        
                     - id: sg-009a95a89ccf7a65b
                     - ingress
                        
                         - 0
                            
                             - cidr_blocks
                                
                                 - 0: 0.0.0.0/0
                                
                             - description: HTTP
                             - from_port: 80
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: tcp
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 80
                            
                         - 1
                            
                             - cidr_blocks
                                
                                
                             - description: Ingress Rule
                             - from_port: 0
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: -1
                             - security_groups
                                
                                
                             - self: true
                             - to_port: 0
                            
                        
                     - name: ibft4_http_sg-20200614141253785600000004
                     - name_prefix: ibft4_http_sg-
                     - owner_id: 133708189211
                     - revoke_rules_on_delete: false
                     - tags
                        
                         - Name: ibft4_http_sg
                        
                     - timeouts: null
                     - vpc_id: vpc-0881077c57fc53714
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6NjAwMDAwMDAwMDAwfSwic2NoZW1hX3ZlcnNpb24iOiIxIn0=
                 - dependencies
                    
                     - 0: module.vpc.aws_vpc.this
                    
                
            
        
     - 133
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group
         - name: this_name_prefix
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 1
                 - attributes
                    
                     - description: ibft4_https_sg
                     - egress
                        
                         - 0
                            
                             - cidr_blocks
                                
                                 - 0: 0.0.0.0/0
                                
                             - description: All protocols
                             - from_port: 0
                             - ipv6_cidr_blocks
                                
                                 - 0: ::/0
                                
                             - prefix_list_ids
                                
                                
                             - protocol: -1
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 0
                            
                        
                     - id: sg-00e1dcb6354d28bdb
                     - ingress
                        
                         - 0
                            
                             - cidr_blocks
                                
                                 - 0: 0.0.0.0/0
                                
                             - description: HTTPS
                             - from_port: 443
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: tcp
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 443
                            
                         - 1
                            
                             - cidr_blocks
                                
                                
                             - description: Ingress Rule
                             - from_port: 0
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: -1
                             - security_groups
                                
                                
                             - self: true
                             - to_port: 0
                            
                        
                     - name: ibft4_https_sg-20200614141253460100000001
                     - name_prefix: ibft4_https_sg-
                     - owner_id: 133708189211
                     - revoke_rules_on_delete: false
                     - tags
                        
                         - Name: ibft4_https_sg
                        
                     - timeouts: null
                     - vpc_id: vpc-0881077c57fc53714
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6NjAwMDAwMDAwMDAwfSwic2NoZW1hX3ZlcnNpb24iOiIxIn0=
                 - dependencies
                    
                     - 0: module.vpc.aws_vpc.this
                    
                
            
        
     - 134
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group
         - name: this_name_prefix
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 1
                 - attributes
                    
                     - description: ibft4_ssh_sg
                     - egress
                        
                         - 0
                            
                             - cidr_blocks
                                
                                 - 0: 0.0.0.0/0
                                
                             - description: All protocols
                             - from_port: 0
                             - ipv6_cidr_blocks
                                
                                 - 0: ::/0
                                
                             - prefix_list_ids
                                
                                
                             - protocol: -1
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 0
                            
                        
                     - id: sg-0fd0fd5fe13b9edca
                     - ingress
                        
                         - 0
                            
                             - cidr_blocks
                                
                                 - 0: 0.0.0.0/0
                                
                             - description: SSH
                             - from_port: 22
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: tcp
                             - security_groups
                                
                                
                             - self: false
                             - to_port: 22
                            
                         - 1
                            
                             - cidr_blocks
                                
                                
                             - description: Ingress Rule
                             - from_port: 0
                             - ipv6_cidr_blocks
                                
                                
                             - prefix_list_ids
                                
                                
                             - protocol: -1
                             - security_groups
                                
                                
                             - self: true
                             - to_port: 0
                            
                        
                     - name: ibft4_ssh_sg-20200614141253784900000003
                     - name_prefix: ibft4_ssh_sg-
                     - owner_id: 133708189211
                     - revoke_rules_on_delete: false
                     - tags
                        
                         - Name: ibft4_ssh_sg
                        
                     - timeouts: null
                     - vpc_id: vpc-0881077c57fc53714
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6NjAwMDAwMDAwMDAwfSwic2NoZW1hX3ZlcnNpb24iOiIxIn0=
                 - dependencies
                    
                     - 0: module.vpc.aws_vpc.this
                    
                
            
        
     - 135
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_egress_rules
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 136
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_egress_rules
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 137
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_egress_rules
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 138
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_egress_with_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 139
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_egress_with_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 140
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_egress_with_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 141
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_egress_with_ipv6_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 142
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_egress_with_ipv6_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 143
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_egress_with_ipv6_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 144
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_egress_with_self
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 145
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_egress_with_self
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 146
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_egress_with_self
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 147
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_egress_with_source_security_group_id
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 148
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_egress_with_source_security_group_id
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 149
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_egress_with_source_security_group_id
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 150
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_ingress_rules
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 151
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_ingress_rules
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 152
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_ingress_rules
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 153
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_ingress_with_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 154
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_ingress_with_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 155
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_ingress_with_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 156
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_ingress_with_ipv6_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 157
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_ingress_with_ipv6_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 158
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_ingress_with_ipv6_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 159
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_ingress_with_self
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 160
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_ingress_with_self
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 161
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_ingress_with_self
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 162
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_ingress_with_source_security_group_id
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 163
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_ingress_with_source_security_group_id
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 164
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: computed_ingress_with_source_security_group_id
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 165
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: egress_rules
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 2
                 - attributes
                    
                     - cidr_blocks
                        
                         - 0: 0.0.0.0/0
                        
                     - description: All protocols
                     - from_port: 0
                     - id: sgrule-1992700788
                     - ipv6_cidr_blocks
                        
                         - 0: ::/0
                        
                     - prefix_list_ids
                        
                        
                     - protocol: -1
                     - security_group_id: sg-009a95a89ccf7a65b
                     - self: false
                     - source_security_group_id: null
                     - to_port: 0
                     - type: egress
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: module.http_80_security_group.module.sg.aws_security_group.this
                     - 1: module.http_80_security_group.module.sg.aws_security_group.this_name_prefix
                     - 2: module.vpc.aws_vpc.this
                    
                
            
        
     - 166
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: egress_rules
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 2
                 - attributes
                    
                     - cidr_blocks
                        
                         - 0: 0.0.0.0/0
                        
                     - description: All protocols
                     - from_port: 0
                     - id: sgrule-3560455182
                     - ipv6_cidr_blocks
                        
                         - 0: ::/0
                        
                     - prefix_list_ids
                        
                        
                     - protocol: -1
                     - security_group_id: sg-00e1dcb6354d28bdb
                     - self: false
                     - source_security_group_id: null
                     - to_port: 0
                     - type: egress
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: module.https_443_security_group.module.sg.aws_security_group.this
                     - 1: module.https_443_security_group.module.sg.aws_security_group.this_name_prefix
                     - 2: module.vpc.aws_vpc.this
                    
                
            
        
     - 167
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: egress_rules
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 2
                 - attributes
                    
                     - cidr_blocks
                        
                         - 0: 0.0.0.0/0
                        
                     - description: All protocols
                     - from_port: 0
                     - id: sgrule-65858084
                     - ipv6_cidr_blocks
                        
                         - 0: ::/0
                        
                     - prefix_list_ids
                        
                        
                     - protocol: -1
                     - security_group_id: sg-0fd0fd5fe13b9edca
                     - self: false
                     - source_security_group_id: null
                     - to_port: 0
                     - type: egress
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: module.ssh_security_group.module.sg.aws_security_group.this
                     - 1: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 2: module.vpc.aws_vpc.this
                    
                
            
        
     - 168
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: egress_with_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 169
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: egress_with_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 170
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: egress_with_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 171
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: egress_with_ipv6_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 172
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: egress_with_ipv6_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 173
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: egress_with_ipv6_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 174
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: egress_with_self
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 175
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: egress_with_self
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 176
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: egress_with_self
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 177
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: egress_with_source_security_group_id
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 178
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: egress_with_source_security_group_id
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 179
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: egress_with_source_security_group_id
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 180
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: ingress_rules
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 2
                 - attributes
                    
                     - cidr_blocks
                        
                         - 0: 0.0.0.0/0
                        
                     - description: HTTP
                     - from_port: 80
                     - id: sgrule-2846444832
                     - ipv6_cidr_blocks
                        
                        
                     - prefix_list_ids
                        
                        
                     - protocol: tcp
                     - security_group_id: sg-009a95a89ccf7a65b
                     - self: false
                     - source_security_group_id: null
                     - to_port: 80
                     - type: ingress
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: module.http_80_security_group.module.sg.aws_security_group.this
                     - 1: module.http_80_security_group.module.sg.aws_security_group.this_name_prefix
                     - 2: module.vpc.aws_vpc.this
                    
                
            
        
     - 181
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: ingress_rules
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 2
                 - attributes
                    
                     - cidr_blocks
                        
                         - 0: 0.0.0.0/0
                        
                     - description: HTTPS
                     - from_port: 443
                     - id: sgrule-4087396957
                     - ipv6_cidr_blocks
                        
                        
                     - prefix_list_ids
                        
                        
                     - protocol: tcp
                     - security_group_id: sg-00e1dcb6354d28bdb
                     - self: false
                     - source_security_group_id: null
                     - to_port: 443
                     - type: ingress
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: module.https_443_security_group.module.sg.aws_security_group.this
                     - 1: module.https_443_security_group.module.sg.aws_security_group.this_name_prefix
                     - 2: module.vpc.aws_vpc.this
                    
                
            
        
     - 182
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: ingress_rules
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 2
                 - attributes
                    
                     - cidr_blocks
                        
                         - 0: 0.0.0.0/0
                        
                     - description: SSH
                     - from_port: 22
                     - id: sgrule-4175333840
                     - ipv6_cidr_blocks
                        
                        
                     - prefix_list_ids
                        
                        
                     - protocol: tcp
                     - security_group_id: sg-0fd0fd5fe13b9edca
                     - self: false
                     - source_security_group_id: null
                     - to_port: 22
                     - type: ingress
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: module.ssh_security_group.module.sg.aws_security_group.this
                     - 1: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 2: module.vpc.aws_vpc.this
                    
                
            
        
     - 183
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: ingress_with_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 184
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: ingress_with_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 185
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: ingress_with_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 186
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: ingress_with_ipv6_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 187
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: ingress_with_ipv6_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 188
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: ingress_with_ipv6_cidr_blocks
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 189
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: ingress_with_self
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 2
                 - attributes
                    
                     - cidr_blocks
                        
                        
                     - description: Ingress Rule
                     - from_port: 0
                     - id: sgrule-691721153
                     - ipv6_cidr_blocks
                        
                        
                     - prefix_list_ids
                        
                        
                     - protocol: -1
                     - security_group_id: sg-009a95a89ccf7a65b
                     - self: true
                     - source_security_group_id: sg-009a95a89ccf7a65b
                     - to_port: 0
                     - type: ingress
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: module.http_80_security_group.module.sg.aws_security_group.this
                     - 1: module.http_80_security_group.module.sg.aws_security_group.this_name_prefix
                     - 2: module.vpc.aws_vpc.this
                    
                
            
        
     - 190
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: ingress_with_self
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 2
                 - attributes
                    
                     - cidr_blocks
                        
                        
                     - description: Ingress Rule
                     - from_port: 0
                     - id: sgrule-3964050827
                     - ipv6_cidr_blocks
                        
                        
                     - prefix_list_ids
                        
                        
                     - protocol: -1
                     - security_group_id: sg-00e1dcb6354d28bdb
                     - self: true
                     - source_security_group_id: sg-00e1dcb6354d28bdb
                     - to_port: 0
                     - type: ingress
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: module.https_443_security_group.module.sg.aws_security_group.this
                     - 1: module.https_443_security_group.module.sg.aws_security_group.this_name_prefix
                     - 2: module.vpc.aws_vpc.this
                    
                
            
        
     - 191
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: ingress_with_self
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 2
                 - attributes
                    
                     - cidr_blocks
                        
                        
                     - description: Ingress Rule
                     - from_port: 0
                     - id: sgrule-1373928833
                     - ipv6_cidr_blocks
                        
                        
                     - prefix_list_ids
                        
                        
                     - protocol: -1
                     - security_group_id: sg-0fd0fd5fe13b9edca
                     - self: true
                     - source_security_group_id: sg-0fd0fd5fe13b9edca
                     - to_port: 0
                     - type: ingress
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjIifQ==
                 - dependencies
                    
                     - 0: module.ssh_security_group.module.sg.aws_security_group.this
                     - 1: module.ssh_security_group.module.sg.aws_security_group.this_name_prefix
                     - 2: module.vpc.aws_vpc.this
                    
                
            
        
     - 192
        
         - module: module.http_80_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: ingress_with_source_security_group_id
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 193
        
         - module: module.https_443_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: ingress_with_source_security_group_id
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 194
        
         - module: module.ssh_security_group.module.sg
         - mode: managed
         - type: aws_security_group_rule
         - name: ingress_with_source_security_group_id
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 195
        
         - module: module.vpc
         - mode: managed
         - type: aws_subnet
         - name: database
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 196
        
         - module: module.vpc
         - mode: managed
         - type: aws_subnet
         - name: elasticache
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 197
        
         - module: module.vpc
         - mode: managed
         - type: aws_subnet
         - name: intra
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 198
        
         - module: module.vpc
         - mode: managed
         - type: aws_subnet
         - name: private
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 1
                 - attributes
                    
                     - assign_ipv6_address_on_creation: false
                     - availability_zone: us-west-1a
                     - availability_zone_id: usw1-az1
                     - cidr_block: 10.0.4.0/24
                     - id: subnet-06e62b6d0c1d8e1b0
                     - ipv6_cidr_block: 
                     - ipv6_cidr_block_association_id: 
                     - map_public_ip_on_launch: false
                     - owner_id: 133708189211
                     - tags
                        
                         - Name: ibft4-private-us-west-1a
                         - terraform: true
                         - vpc: ibft4
                        
                     - timeouts: null
                     - vpc_id: vpc-0881077c57fc53714
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6MTIwMDAwMDAwMDAwMH0sInNjaGVtYV92ZXJzaW9uIjoiMSJ9
                 - dependencies
                    
                     - 0: module.vpc.aws_vpc.this
                     - 1: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
             - 1
                
                 - index_key: 1
                 - schema_version: 1
                 - attributes
                    
                     - assign_ipv6_address_on_creation: false
                     - availability_zone: us-west-1c
                     - availability_zone_id: usw1-az3
                     - cidr_block: 10.0.5.0/24
                     - id: subnet-003fdf10e9015f50a
                     - ipv6_cidr_block: 
                     - ipv6_cidr_block_association_id: 
                     - map_public_ip_on_launch: false
                     - owner_id: 133708189211
                     - tags
                        
                         - Name: ibft4-private-us-west-1c
                         - terraform: true
                         - vpc: ibft4
                        
                     - timeouts: null
                     - vpc_id: vpc-0881077c57fc53714
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6MTIwMDAwMDAwMDAwMH0sInNjaGVtYV92ZXJzaW9uIjoiMSJ9
                 - dependencies
                    
                     - 0: module.vpc.aws_vpc.this
                     - 1: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
             - 2
                
                 - index_key: 2
                 - schema_version: 1
                 - attributes
                    
                     - assign_ipv6_address_on_creation: false
                     - availability_zone: us-west-1a
                     - availability_zone_id: usw1-az1
                     - cidr_block: 10.0.6.0/24
                     - id: subnet-0d97d5bad724ebcac
                     - ipv6_cidr_block: 
                     - ipv6_cidr_block_association_id: 
                     - map_public_ip_on_launch: false
                     - owner_id: 133708189211
                     - tags
                        
                         - Name: ibft4-private-us-west-1a
                         - terraform: true
                         - vpc: ibft4
                        
                     - timeouts: null
                     - vpc_id: vpc-0881077c57fc53714
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6MTIwMDAwMDAwMDAwMH0sInNjaGVtYV92ZXJzaW9uIjoiMSJ9
                 - dependencies
                    
                     - 0: module.vpc.aws_vpc.this
                     - 1: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
            
        
     - 199
        
         - module: module.vpc
         - mode: managed
         - type: aws_subnet
         - name: public
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 1
                 - attributes
                    
                     - assign_ipv6_address_on_creation: false
                     - availability_zone: us-west-1a
                     - availability_zone_id: usw1-az1
                     - cidr_block: 10.0.1.0/24
                     - id: subnet-07df858a84da18c00
                     - ipv6_cidr_block: 
                     - ipv6_cidr_block_association_id: 
                     - map_public_ip_on_launch: true
                     - owner_id: 133708189211
                     - tags
                        
                         - Name: ibft4-public-us-west-1a
                         - terraform: true
                         - vpc: ibft4
                        
                     - timeouts: null
                     - vpc_id: vpc-0881077c57fc53714
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6MTIwMDAwMDAwMDAwMH0sInNjaGVtYV92ZXJzaW9uIjoiMSJ9
                 - dependencies
                    
                     - 0: module.vpc.aws_vpc.this
                     - 1: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
             - 1
                
                 - index_key: 1
                 - schema_version: 1
                 - attributes
                    
                     - assign_ipv6_address_on_creation: false
                     - availability_zone: us-west-1c
                     - availability_zone_id: usw1-az3
                     - cidr_block: 10.0.2.0/24
                     - id: subnet-0d619244f43aeea5e
                     - ipv6_cidr_block: 
                     - ipv6_cidr_block_association_id: 
                     - map_public_ip_on_launch: true
                     - owner_id: 133708189211
                     - tags
                        
                         - Name: ibft4-public-us-west-1c
                         - terraform: true
                         - vpc: ibft4
                        
                     - timeouts: null
                     - vpc_id: vpc-0881077c57fc53714
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6MTIwMDAwMDAwMDAwMH0sInNjaGVtYV92ZXJzaW9uIjoiMSJ9
                 - dependencies
                    
                     - 0: module.vpc.aws_vpc.this
                     - 1: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
             - 2
                
                 - index_key: 2
                 - schema_version: 1
                 - attributes
                    
                     - assign_ipv6_address_on_creation: false
                     - availability_zone: us-west-1a
                     - availability_zone_id: usw1-az1
                     - cidr_block: 10.0.3.0/24
                     - id: subnet-0005aeeaff667e9f9
                     - ipv6_cidr_block: 
                     - ipv6_cidr_block_association_id: 
                     - map_public_ip_on_launch: true
                     - owner_id: 133708189211
                     - tags
                        
                         - Name: ibft4-public-us-west-1a
                         - terraform: true
                         - vpc: ibft4
                        
                     - timeouts: null
                     - vpc_id: vpc-0881077c57fc53714
                    
                 - private: eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjo2MDAwMDAwMDAwMDAsImRlbGV0ZSI6MTIwMDAwMDAwMDAwMH0sInNjaGVtYV92ZXJzaW9uIjoiMSJ9
                 - dependencies
                    
                     - 0: module.vpc.aws_vpc.this
                     - 1: module.vpc.aws_vpc_ipv4_cidr_block_association.this
                    
                
            
        
     - 200
        
         - module: module.vpc
         - mode: managed
         - type: aws_subnet
         - name: redshift
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 201
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc
         - name: this
         - each: list
         - provider: provider.aws
         - instances
            
             - 0
                
                 - index_key: 0
                 - schema_version: 1
                 - attributes
                    
                     - assign_generated_ipv6_cidr_block: false
                     - cidr_block: 10.0.0.0/16
                     - default_network_acl_id: acl-09757189f81464554
                     - default_route_table_id: rtb-03538f05e18540a5e
                     - default_security_group_id: sg-02ebdfa0219604058
                     - dhcp_options_id: dopt-84676ae3
                     - enable_classiclink: false
                     - enable_classiclink_dns_support: false
                     - enable_dns_hostnames: true
                     - enable_dns_support: true
                     - id: vpc-0881077c57fc53714
                     - instance_tenancy: default
                     - ipv6_association_id: 
                     - ipv6_cidr_block: 
                     - main_route_table_id: rtb-03538f05e18540a5e
                     - owner_id: 133708189211
                     - tags
                        
                         - Name: ibft4
                         - terraform: true
                         - vpc: ibft4
                        
                    
                 - private: eyJzY2hlbWFfdmVyc2lvbiI6IjEifQ==
                
            
        
     - 202
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_dhcp_options
         - name: this
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 203
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_dhcp_options_association
         - name: this
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 204
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: access_analyzer
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 205
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: acm_pca
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 206
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: apigw
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 207
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: appmesh_envoy_management
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 208
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: appstream
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 209
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: athena
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 210
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: auto_scaling_plans
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 211
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: cloud_directory
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 212
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: cloudformation
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 213
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: cloudtrail
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 214
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: codebuild
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 215
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: codecommit
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 216
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: codepipeline
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 217
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: config
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 218
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: datasync
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 219
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: dynamodb
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 220
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: ebs
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 221
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: ec2
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 222
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: ec2_autoscaling
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 223
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: ec2messages
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 224
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: ecr_api
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 225
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: ecr_dkr
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 226
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: ecs
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 227
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: ecs_agent
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 228
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: ecs_telemetry
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 229
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: efs
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 230
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: elastic_inference_runtime
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 231
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: elasticbeanstalk
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 232
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: elasticbeanstalk_health
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 233
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: elasticloadbalancing
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 234
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: emr
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 235
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: events
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 236
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: git_codecommit
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 237
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: glue
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 238
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: kinesis_firehose
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 239
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: kinesis_streams
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 240
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: kms
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 241
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: logs
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 242
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: monitoring
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 243
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: qldb_session
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 244
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: rekognition
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 245
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: s3
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 246
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: sagemaker_api
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 247
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: sagemaker_notebook
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 248
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: sagemaker_runtime
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 249
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: secretsmanager
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 250
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: servicecatalog
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 251
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: ses
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 252
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: sms
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 253
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: sns
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 254
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: sqs
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 255
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: ssm
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 256
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: ssmmessages
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 257
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: states
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 258
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: storagegateway
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 259
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: sts
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 260
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: transfer
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 261
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: transferserver
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 262
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint
         - name: workspaces
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 263
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint_route_table_association
         - name: intra_dynamodb
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 264
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint_route_table_association
         - name: intra_s3
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 265
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint_route_table_association
         - name: private_dynamodb
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 266
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint_route_table_association
         - name: private_s3
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 267
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint_route_table_association
         - name: public_dynamodb
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 268
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_endpoint_route_table_association
         - name: public_s3
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 269
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpc_ipv4_cidr_block_association
         - name: this
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 270
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpn_gateway
         - name: this
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 271
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpn_gateway_attachment
         - name: this
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 272
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpn_gateway_route_propagation
         - name: intra
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 273
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpn_gateway_route_propagation
         - name: private
         - each: list
         - provider: provider.aws
         - instances
            
            
        
     - 274
        
         - module: module.vpc
         - mode: managed
         - type: aws_vpn_gateway_route_propagation
         - name: public
         - each: list
         - provider: provider.aws
         - instances
            
            
        
    
