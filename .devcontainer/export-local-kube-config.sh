#!/usr/bin/env bash
# Setting the namespace preference
namespace="default"
if [[ $KUBERNETES_NAMESPACE ]]; then
  namespace=${KUBERNETES_NAMESPACE}
fi

kubectl config set-context --current --namespace=${namespace}