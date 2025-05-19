# poc_kc
1. Ensure you have Docker Desktop installed
2. Enable Docker swarm by doing "docker swarm init"
3. To stop and remove docker swarm instances do "docker swarm leave --force" 
4. In root folder  run below commands 
   docker build --no-cache -t keycloak-custom keycloak   
   docker build --no-cache -t admin-app admin-app
5. docker stack deploy -c docker-compose.yml keycloak-stack
6. You will see output
      Creating network keycloak-stack_default
      Creating service keycloak-stack_keycloak
      Creating service keycloak-stack_postgres
      Creating service keycloak-stack_admin-app

   After finish the start up in 10-20  seconds Keycloak admin will be availble on localhost:8080

   Keycloak is pre configured with below 

     Two realms will be pre-created
       REALM 1 : PAccount (realm-account.json )
         Users
             barun (admin role in PAccount realm )
             shialendra (admin role in PAccount realm )
         Client
             dummyclient
                 Direct grant access enabled

   REALM 2 :  EndUsers (realm-endUsers.json )
       Users
           User1 (  no role assigned  )
           User2  ( no role assigned )
       Client
           endusers-client 
               Direct grant access enabled 

   GO LANG Admin app POC is being evolved (you can contribute )
   Go lang admin app available on localhost:9081/login.html 

  NOTE : Though postgraceSQL container is coming up in swarm

  HOW to Start development 
  Using Dockerfile bring up the docker containers and dependenies  
  Run the keycloak or other containers without swarm so that you can see the logs of keycloak and other containers . 

