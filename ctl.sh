#!/bin/bash -e

IMAGE_NAME=registry.gitlab.com/demo_proj:latest
DOCKER_AUTH_TOKEN=xxxxx

usage() {
  cat <<EOF
Usage:
  sh ./$(basename $0) --haproxy  update haproxy config

  sh ./$(basename $0) --migrate  run sql migrations
  sh ./$(basename $0) --run      build and run application
  sh ./$(basename $0) --prod     deploy to prod
EOF
  exit 1
}

log() {
  local DATETIME=`date '+%Y/%m/%d %H:%M:%S'`
  echo "${DATETIME} ${1}"
}


export_ldflags() {
	build_date=$(date +"%d.%m_%H:%M")
	git_rev=$(git rev-parse --short HEAD)
  
	export LDFLAGS="-X main.AppVersion=${build_date}_${git_rev}"

  echo "${LDFLAGS}"
}

migrate() {
  log "Run migrations"
  goose -dir="./db" postgres "host=${DATABASE_HOST} port=${DATABASE_PORT} user=${DATABASE_USER} password=${DATABASE_PASSWORD} database=${DATABASE_NAME} sslmode=${DATABASE_SSL_MODE} sslrootcert=${DATABASE_SSL_CERT}" up -v
}

build() {
  log "Run Make build: [${IMAGE_NAME}]"
	export_ldflags
  docker build -t ${IMAGE_NAME} . --build-arg LDFLAGS="${LDFLAGS}"
}

push() {
  log "Push docker image: [${IMAGE_NAME}]"
  docker login --username oauth --password $DOCKER_AUTH_TOKEN registry.gitlab.com
  docker push ${IMAGE_NAME}
}

haproxy() {
  sh ../../infra/configuration/ctl.sh --haproxy
}

run_local() {
  clear
  CURRENT_HOST_NAME=$(scutil --get LocalHostName)
  log "Build and run application at $CURRENT_HOST_NAME. Please wait.."
  go run cmd/api/*
}

case "$1" in
  "--build")
    build
    ;;
  "--docker")
    build
    run_docker
    ;;
  "--migrate")
    migrate
    ;;
  "--push")
    build
    push
    ;;
  "--ldf")
    export_ldflags
    ;;
  "--run")
    run_local
    ;;
  *)
    usage
    ;;
esac
