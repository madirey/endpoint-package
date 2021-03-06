package install

const snapshotYml = `version: '2.3'
services:
  elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:${STACK_VERSION}
    healthcheck:
      test: ["CMD", "curl", "-f", "-u", "elastic:changeme", "http://127.0.0.1:9200/"]
      retries: 300
      interval: 1s
    environment:
    - "ES_JAVA_OPTS=-Xms1g -Xmx1g"
    - "network.host="
    - "transport.host=127.0.0.1"
    - "http.host=0.0.0.0"
    - "indices.id_field_data.enabled=true"
    - "xpack.license.self_generated.type=trial"
    - "xpack.security.enabled=true"
    - "xpack.security.authc.api_key.enabled=true"
    - "ELASTIC_PASSWORD=changeme"
    ports:
      - "127.0.0.1:9200:9200"

  kibana:
    image: docker.elastic.co/kibana/kibana:${STACK_VERSION}
    depends_on:
      elasticsearch:
        condition: service_healthy
      package-registry:
        condition: service_healthy
    healthcheck:
      test: "curl -f http://127.0.0.1:5601/login | grep kbn-injected-metadata 2>&1 >/dev/null"
      retries: 600
      interval: 1s
    volumes:
      - ./kibana.config.yml:/usr/share/kibana/config/kibana.yml
    ports:
      - "127.0.0.1:5601:5601"

  package-registry:
    build:
      context: .
      dockerfile: Dockerfile.package-registry
    healthcheck:
      test: ["CMD", "curl", "-f", "http://127.0.0.1:8080"]
      retries: 300
      interval: 1s
    ports:
      - "127.0.0.1:8080:8080"

  is_ready:
    image: tianon/true
    depends_on:
      kibana:
        condition: service_healthy
`
