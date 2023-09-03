local_resource(
  'courses-api-compile',
  cmd='cd courses-api; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/courses-api',
  deps=['courses-api/main.go'],
  labels=['courses-api'],
)

local_resource(
  'courses-frontend-compile',
  cmd='cd courses-frontend; npm run build',
  deps=[
    'courses-frontend/src/',
    'courses-frontend/package.json',
    'courses-frontend/package-lock.json',
    'courses-frontend/public',
  ],
  labels=['courses-frontend'],
)

docker_build(
  'courses-api',
  'courses-api/',
  dockerfile='./courses-api/Dockerfile',
  live_update=[
    sync('./courses-api/build/courses-api', '/courses-api'),
  ]
)

docker_build(
  'courses-frontend',
  'courses-frontend/',
  dockerfile='./courses-frontend/Dockerfile',
  live_update=[
    sync('./courses-frontend/build/*', '/usr/share/nginx/html/'),
  ]
)

# docker_compose(
# './docker-compose.yaml'
# )

