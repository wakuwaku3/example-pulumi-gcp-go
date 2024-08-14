FROM mcr.microsoft.com/vscode/devcontainers/go:1-bookworm

RUN apt update && \
  apt install -y \
  apt-transport-https \
  ca-certificates \
  curl \
  gnupg

# google-cloud-cli
RUN curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | gpg --dearmor -o /usr/share/keyrings/cloud.google.gpg
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
RUN apt update && apt install -y google-cloud-cli google-cloud-sdk-gke-gcloud-auth-plugin

WORKDIR /home/vscode
USER vscode

# pulumi
RUN curl -fsSL https://get.pulumi.com | sh
ENV PATH="/home/vscode/.pulumi/bin:$PATH"
