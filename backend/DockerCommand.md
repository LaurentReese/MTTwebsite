tuto : https://tutorialedge.net/golang/go-docker-tutorial/#why-docker

from the directory where Dockerfile is, run the command :

==> sudo docker build . -t myapp

see the results :

==> sudo docker images

run it 

-==> sudo docker run -p 8090:8090 -it myapp
