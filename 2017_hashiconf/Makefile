.PHONY: graphs
docker_image=2017_hashiconf
container_name=hashiconf
pwd=$(shell pwd)

container:
	docker build --no-cache -t $(docker_image) ./

present: graphs
	$(GOPATH)/bin/present -http 0.0.0.0:3999 -notes ./puppet_overview.slide

down:
	@docker rm -f $(container_name) 2>/dev/null || true

# Brings up the dev instance of mf2.
up: down
	@chmod 0755 run.sh
	docker run -d \
		-v $(pwd):/mnt/build \
		-v $(pwd)/run.sh:/run.sh \
		-p 3999:3999 \
		--name $(container_name) \
		$(docker_image)
	docker logs -f $(container_name)

console:
	docker exec -t -i $(container_name) /bin/bash

graphs:
	mkdir -p images/{graphs,sequences,diagrams}
	java -jar ~/plantuml.jar -o $(pwd)/images/diagrams/ graphs/original.uml
	java -jar ~/plantuml.jar -o $(pwd)/images/diagrams/ graphs/original_singlenode.uml
	java -jar ~/plantuml.jar -o $(pwd)/images/diagrams/ graphs/with_nomad.uml
	convert images/diagrams/original_singlenode.png images/diagrams/original_singlenode.png images/diagrams/original_singlenode.png -append images/diagrams/multi_customers.png
