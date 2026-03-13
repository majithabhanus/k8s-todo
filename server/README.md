# cmd for run the mysql container with env variables
  docker run --name mysql-container \
  --env-file .env \
  -e MYSQL_DATABASE=$DB_NAME \
  -e MYSQL_USER=$DB_USER \
  -e MYSQL_PASSWORD=$DB_PASS \
  -p 3306:3306 \
  -d mysql:8.0

# go to mysql container bash in mydb
  mysql -h 127.0.0.1 -P 3306 -u admin -p mydb

# cmd for run go server
  go run cmd/main.go  

# cmd for run the mysql container
  docker run --name mysql-container \
  --network todo-net \
  -e MYSQL_ROOT_PASSWORD=root \
  -e MYSQL_DATABASE=mydb \
  -e MYSQL_USER=admin \
  -e MYSQL_PASSWORD=admin123 \
  -p 3306:3306 \
  -d mysql:8.0

# cmd for run the todo app container  
   
  docker rm -f todo-container

docker run --name todo-container \
  --network todo-net \
  --env DB_USER=admin \
  --env DB_PASS=admin123 \
  --env DB_HOST=mysql-container \
  --env DB_PORT=3306 \
  --env DB_NAME=mydb \
  --env JWT_SECRET='G9s!zX7$eA2@tKp' \
  -p 9000:9000 \
  go-todo
 

 # cmd or create docker network
   
  docker network create  todo-net

# print env on todo-container
  docker exec -it todo-container printenv




# port forward
  kubectl port-forward <target> [HOST_PORT]:[POD_PORT]
  
Left (host port) → where you connect 

Right (pod port) → where the app listens inside K8s  


| Role          | Path              | Description                      |
| ------------- | ----------------- | -------------------------------- |
| **hostPath**  | `/mnt/data/mysql` | Source path on the host node     |
| **mountPath** | `/var/lib/mysql`  | Target path inside the container |




| Environment | NodePort Works?  | Port-Forward? | LoadBalancer? | Ingress Support |
| ----------- | ---------------- | ------------- | ------------- | --------------- |
| Kind        | ❌ Not by default | ✅ Yes         | ❌ No          | ✅ Yes (manual)  |
| Minikube    | ✅ Yes            | ✅ Yes         | ❌ (fake LB)   | ✅ Yes           |
| Cloud (GKE) | ✅ Yes            | ✅ (via SSH)   | ✅ Yes         | ✅ Yes           |


❓Why Doesn’t ClusterIP Work for External Access?

Because:

ClusterIP services are only routable inside the Kubernetes cluster.

The IP it assigns (like 10.96.12.3) is from a virtual network internal to the cluster.

Your browser or curl from your laptop can’t reach that virtual network.


| Service Type   | What Port You Use in URL    | Publicly Accessible?  | Notes                                     |
| -------------- | --------------------------- | --------------------- | ----------------------------------------- |
| `NodePort`     | `http://<NodeIP>:nodePort`  | ✅ Yes (with firewall) | You must use the `nodePort` (e.g., 32625) |
| `LoadBalancer` | `http://<EXTERNAL-IP>:port` | ✅ Yes                 | You use the `port`, not the nodePort      |
| `ClusterIP`    | ❌ Not accessible externally | ❌ No                  | For internal cluster use only             |




# KOPS

