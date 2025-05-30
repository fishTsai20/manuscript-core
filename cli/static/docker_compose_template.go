package static

var DockerComposeTemplate = `
name: {{.Name}}
services:
  jobmanager:
    image: repository.chainbase.com/manuscript-node/manuscript-node:v1.3.0
    user: "manuscript"
    environment:
      - |
        FLINK_PROPERTIES=
        jobmanager.rpc.address: jobmanager
    command: "standalone-job --job-classname com.chainbase.manuscript.ETLProcessor file:///opt/flink/manuscript.yaml --fromSavepoint /opt/flink/savepoint"
    ports:
      - "{{.Port}}:8081"
    volumes:
      - {{ .CkDir }}:/opt/flink/checkpoint
      - {{ .SpDir }}:/opt/flink/savepoint
      - {{ .LogDir }}:/opt/flink/log
      - ./manuscript.yaml:/opt/flink/manuscript.yaml
    networks:
      - ms_network_{{ .Name }}

  taskmanager:
    image: repository.chainbase.com/manuscript-node/manuscript-node:v1.3.0
    user: "manuscript"
    environment:
      - |
        FLINK_PROPERTIES=
        jobmanager.rpc.address: jobmanager
    depends_on:
      - jobmanager
    command: "taskmanager"
    scale: 1
    volumes:
      - {{ .CkDir }}:/opt/flink/checkpoint
      - {{ .SpDir }}:/opt/flink/savepoint
      - {{ .LogDir }}:/opt/flink/log
      - ./manuscript.yaml:/opt/flink/manuscript.yaml
    networks:
      - ms_network_{{ .Name }}

networks:
  ms_network_{{ .Name }}:`

var DockerComposeWithPostgresqlContent = `
name: {{.Name}}
services:
  jobmanager:
    image: repository.chainbase.com/manuscript-node/manuscript-node:v1.3.0
    user: "manuscript"
    environment:
      - |
        FLINK_PROPERTIES=
        jobmanager.rpc.address: jobmanager
    command: "standalone-job --job-classname com.chainbase.manuscript.ETLProcessor file:///opt/flink/manuscript.yaml --fromSavepoint /opt/flink/savepoint"
    ports:
      - "{{.Port}}:8081"
    depends_on:
      postgres:
        condition: service_healthy
    volumes:
      - {{ .CkDir }}:/opt/flink/checkpoint
      - {{ .SpDir }}:/opt/flink/savepoint
      - {{ .LogDir }}:/opt/flink/log
      - ./manuscript.yaml:/opt/flink/manuscript.yaml
    networks:
      - ms_network_{{ .Name }}

  taskmanager:
    image: repository.chainbase.com/manuscript-node/manuscript-node:v1.3.0
    user: "manuscript"
    environment:
      - |
        FLINK_PROPERTIES=
        jobmanager.rpc.address: jobmanager
    depends_on:
      - jobmanager
    command: "taskmanager"
    scale: 1
    volumes:
      - {{ .CkDir }}:/opt/flink/checkpoint
      - {{ .SpDir }}:/opt/flink/savepoint
      - {{ .LogDir }}:/opt/flink/log
      - ./manuscript.yaml:/opt/flink/manuscript.yaml
    networks:
      - ms_network_{{ .Name }}

  postgres:
    image: postgres:16.4
    ports:
      - "{{.DbPort}}:5432"
    volumes:
      - ./postgres_data:/var/lib/postgresql/data
    environment:
      - POSTGRES_PASSWORD=${POSTGRES_PASSWORD:-postgres}
      - POSTGRES_USER=${POSTGRES_USER:-postgres}
      - POSTGRES_DB=${POSTGRES_DB:-public}
    networks:
      - ms_network_{{ .Name }}
    restart: unless-stopped
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U ${POSTGRES_USER:-postgres}"]
      interval: 10s
      timeout: 5s
      retries: 5

  hasura:
    image: {{.GraphQLImage}}
    ports:
      - "{{.GraphQLPort}}:8080"
    depends_on:
      - taskmanager
    environment:
      HASURA_GRAPHQL_DATABASE_URL: postgres://postgres:${POSTGRES_PASSWORD:-postgres}@postgres:5432/{{.Database}}
      HASURA_GRAPHQL_METADATA_DATABASE_URL: postgres://postgres:${POSTGRES_PASSWORD:-postgres}@postgres:5432/{{.Database}}
      HASURA_GRAPHQL_ENABLE_CONSOLE: "true"
    networks:
      - ms_network_{{ .Name }}
    restart: unless-stopped

networks:
  ms_network_{{ .Name }}:`
