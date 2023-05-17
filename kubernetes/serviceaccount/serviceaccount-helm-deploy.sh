#!/bin/bash

namespace=dev
serviceaccount=helm
role=deploy

if [ "${1:-}" == "delete" ]
then
  kubectl delete role "${role}" -n ${namespace}
  kubectl delete rolebinding "${serviceaccount}-${role}" -n ${namespace}
  kubectl delete serviceaccount "${serviceaccount}" -n ${namespace}
  exit
fi


deployRole() {
  kubectl apply -n ${namespace} -f - <<EOF
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: ${role}
  namespace: ${namespace}
rules:
- apiGroups: ["","apps","networking.k8s.io","batch"]
  resources: ["secrets","serviceaccounts","services","deployments","ingresses","daemonsets","jobs"]
  verbs: ["list","get","create","update","patch"]
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["delete"]
EOF
}

kubectl get role "${role}" -n ${namespace} > /dev/null 2>&1
if [ "$?" -eq "1" ]
then
  echo "### Create role"
else
  echo "### Reset role"
fi
deployRole

kubectl get namespace "${namespace}" > /dev/null 2>&1
if [ "$?" -eq "1" ]
then
  echo "### Create namespace"
  kubectl create namespace "${namespace}"
fi

kubectl get rolebinding "${serviceaccount}-${role}" -n ${namespace} > /dev/null 2>&1
if [ "$?" -eq "1" ]
then
  echo "### Create rolebinding"
  kubectl create rolebinding "${serviceaccount}-${role}" \
    --role ${role} \
    --serviceaccount ${namespace}:${serviceaccount} \
    --namespace ${namespace}
fi

kubectl get serviceaccount "${serviceaccount}" -n ${namespace} > /dev/null 2>&1
if [ "$?" -eq "1" ]
then
  echo "### Create serviceaccount"
  kubectl create serviceaccount ${serviceaccount} \
    --namespace ${namespace}

   kubectl apply -n ${namespace} -f - <<EOF
apiVersion: v1
kind: Secret
metadata:
  name: ${serviceaccount}-token
  namespace: ${namespace}
  annotations:
    kubernetes.io/service-account.name: ${serviceaccount}
type: kubernetes.io/service-account-token
EOF

fi

