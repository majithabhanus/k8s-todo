# install necessary tools for eks cluster management then delete 

# install aws cli

# Download, unzip, and install AWS CLI v2
apt install unzip
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install

# verify
aws --version

# configure aws cli (if not already configured)
# aws configure
 

# install kubectl

# Download latest stable kubectl binary
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x kubectl
sudo mv kubectl /usr/local/bin/

# verify
kubectl version --client --short




# install eksctl

# Download latest eksctl tarball and install
curl --silent --location "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
sudo mv /tmp/eksctl /usr/local/bin/

# verify
eksctl version


# to create eks cluster command
 
 #eksctl create cluster --name my-eks-cluster --region $AWS_REGION --nodes 2 --node-type t3.medium --managed


# set region (example: Stockholm eu-north-1)
export AWS_REGION=eu-north-1
aws configure set region $AWS_REGION     # run aws configure first to set credentials/region if not already done

# create cluster with 2 managed nodes (t3.medium)
eksctl create cluster \
  --name my-eks-cluster \
  --region $AWS_REGION \
  --nodes 2 \
  --nodes-min 1 \
  --nodes-max 3 \
  --node-type t3.medium \
  --managed


# to delete eks cluster command

eksctl delete cluster --name my-eks-cluster --region $AWS_REGION


# this command should be run on one ec2 instance only to avoid multiple cluster creations for client apps

# the create another server for jenkins and other ci/cd tools after cluster is created
    ✅ Next step for you:

    Launch 1 EC2 instance (t3.medium)

    Install Jenkins, Docker, kubectl, AWS CLI

    Configure AWS, Github and Docker credentials on Jenkins global credentials

    Create a Jenkins pipeline job with the above Jenkinsfile

    Run it → see your Go app deployed to EKS


    #after install jenkins with these tools install the following plugins
     - AWS Credentials
     - stage view
     - Docker Pipeline


     #set these credentials in jenkins global credentials
      - AWS credentials (access key id and secret access key)
      - Dockerhub credentials (dockerhub username and password)
      - GitHub credentials (github username and personal access token)
    
      #for github personal access token generate it from github settings with repo access



