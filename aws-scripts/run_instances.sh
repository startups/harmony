#!/bin/bash -x
echo "Run Instances" >> tmplog
cd /home/ec2-user/projects/src/harmony-benchmark
./deploy_one_instance.sh global_nodes.txt