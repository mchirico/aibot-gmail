# Build the manager binary
FROM golang:1.15 as builder



WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download


COPY . /workspace/
RUN go get -v -t -d ./...

# Build
# -tags timetzdata
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -tags timetzdata -a -o project main.go
RUN cd  pkg/etcdserver && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -tags timetzdata -a -o server server.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/project .
COPY --from=builder /workspace/pkg/etcdserver/server .


COPY --from=builder --chown=nonroot:nonroot /workspace/credentials /credentials
COPY --from=builder --chown=nonroot:nonroot /workspace/certs /certs


USER nonroot:nonroot

ARG buildtime_variable
ENV GITHUB_TOKEN=$buildtime_variable

ARG webhook
ENV GITHUB_WEBHOOK_SECRET=$webhook

ARG pubsubtoken
ENV PUBSUBTOKEN=$pubsubtoken

ARG build_task_key
ENV TASK_KEY=$build_task_key



ENTRYPOINT ["/project"]
