# docker images | grep golang | awk '{$1="";$2="";$4="";$5="";$6="";$7="";$8="";$9="";$10="";print}'
# to be executed when we are in the MTTwebsite directory

# TO DO : verify that we are in the MTTwebsite directory
set -x # verbose mode in order to follow what is happening

docker rmi 32681733/mtt-frontend --force
docker rmi 32681733/mtt-backend --force

docker images purge
docker images

cd frontend
cat .env.production.local

docker build . -t 32681733/mtt-frontend
docker push 32681733/mtt-frontend

cd ../backend
./before_compile_backend
docker build . -t 32681733/mtt-backend
./after_compile_backend
docker push 32681733/mtt-backend

cd ..
sloppy delete mtt-habitat
sloppy show
sloppy start --var=domain:mtt-habitat.sloppy.zone mtt-sloppy.json
sloppy show mtt-habitat

set +x
echo "sloppy show mtt-habitat"
