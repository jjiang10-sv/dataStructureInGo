#-------------------------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See https://go.microsoft.com/fwlink/?linkid=2090316 for license information.
#-------------------------------------------------------------------------------------------------------------

FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

    # Install kubectx
RUN git clone https://github.com/ahmetb/kubectx /opt/kubectx \
&& ln -s /opt/kubectx/kubectx /usr/local/bin/kubectx \
&& ln -s /opt/kubectx/kubens /usr/local/bin/kubens

# Install kubectl
# Note: Latest version of kubectl may be found at:
# https://github.com/kubernetes/kubernetes/releases
ARG KUBE_LATEST_VERSION="v1.19.2"
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/${KUBE_LATEST_VERSION}/bin/linux/amd64/kubectl
RUN chmod +x ./kubectl
RUN mv ./kubectl /usr/bin/kubectl

# Install helm
# Note: Latest version of helm may be found at
# https://github.com/kubernetes/helm/releases
ARG HELM_VERSION="v3.7.2"
RUN wget -q https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz -O - | tar -xzO linux-amd64/helm > /usr/local/bin/helm \
&& chmod +x /usr/local/bin/helm