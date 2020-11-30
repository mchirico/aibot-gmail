
docker-build:
	docker build --build-arg build_task_key=${TASK_KEY} --build-arg buildtime_variable=${GITHUB_TOKEN} --build-arg webhook=${GITHUB_WEBHOOK_SECRET} --no-cache -t gcr.io/mchirico/aibot:test -f Dockerfile .

push:
	docker push gcr.io/mchirico/aibot:test

build:
	go build -v .

run:
	docker run --name aibot --rm -it -p 3000:3000  gcr.io/mchirico/aibot:test


deploy:
	docker build --build-arg build_task_key=${TASK_KEY} --build-arg buildtime_variable=${GITHUB_TOKEN} --build-arg webhook=${GITHUB_WEBHOOK_SECRET} --build-arg pubsubtoken=${PUBSUBTOKEN} --no-cache -t gcr.io/mchirico/aibot:test -f Dockerfile .
	docker push gcr.io/mchirico/aibot:test
	gcloud beta run deploy aibot  --image gcr.io/mchirico/aibot:test --platform managed \
            --allow-unauthenticated --project mchirico \
            --vpc-connector=conn2 \
            --vpc-egress=all \
            --region us-east1 --port 3000 --max-instances 1  --memory 256Mi



