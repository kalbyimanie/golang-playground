## Prerequisites
- docker-deskop (or your chosen docker tools)
- kubernetes local (or your chosen local kubernetes)
- setup local registry for faster image pull
- skaffold for local CI/CD (for faster local implementation)

## [1] Setup local Image registry
#### 1.0 running local docker registry
`docker run -d -p 5000:5000 --name registry registry:2.7`

#### 1.1 push image to local registry (manual push)

`docker tag <your_image_tag> localhost:5000/<your_image_tag>`

## [2] Setup Local CI/CD

#### 2.0 set your kubecontext to your local kubernetes or docker-desktop
`kubectl config use-context docker-desktop`

#### 2.1 run skaffold dev for automatic build and deploy on changes
`skaffold dev --namespace=my-app`

#### 2.1.1 run skaffold build automatic build only

`skaffold build -p local-registry`


## [3] Expose your app by port-forwarding to your host network
`kubectl port-forward svc/my-app -n my-app 8080:8080`

## [4] Access your app with your chosen browser and fill the url bar with the following address
`http://localhost:8080`
