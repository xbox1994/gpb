package template

import (
	"errors"
	"fmt"
	"os"
)

var k8sConfig = &Template{
	Mode:     Create,
	FilePath: `wps.eks.app.deployment.yaml`,
	Content: `kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: ${app_name}
  namespace: ${app_namespace}
  labels:
    k8s-app: ${app_name}
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: ${app_name}
  template:
    metadata:
      labels:
        k8s-app: ${app_name}
        name: ${app_name}
    spec:
      containers:
      - image: ${image_name}
        name: ${app_name}
        ports:
        - name: backend-http
          containerPort: 8080
        env:
        - name: ENV_MODE
          value: ${env_mode}
---
kind: Service
apiVersion: v1
metadata:
  name: ${app_name}-service
  namespace: ${app_namespace}
  labels:
    k8s-app: ${app_name}
  annotations:
    # Note that the backend talks over HTTP.
    service.beta.kubernetes.io/aws-load-balancer-backend-protocol: http
    # TODO: Fill in with the ARN of your certificate.
    # service.beta.kubernetes.io/aws-load-balancer-ssl-cert: arn:aws:acm:{region}:{user id}:certificate/{id}
    # Only run SSL on the port named "https" below.
    service.beta.kubernetes.io/aws-load-balancer-ssl-ports: "https"
spec:
  selector:
    k8s-app: ${app_name}
  ports:
  - port: 80
    targetPort: backend-http
    name: http
{%- if IS_PROD %}
  - port: 443
    targetPort: backend-http
    name: https
  type: LoadBalancer
{%- else %}
  type: NodePort
{%- endif %}`,
	StdOut: createK8sConfig,
}

func init() {
	AvailableTemplates = append(AvailableTemplates, k8sConfig)
}

func createK8sConfig(template *Template, args ...string) (err error) {
	if len(args) < 2 {
		err = errors.New(`params error`)
	} else {
		projectPath := args[0]
		projectName := args[1]
		absPath := fmt.Sprintf("%s/%s/", projectPath, projectName)
		if file, err1 := os.Create(absPath + template.FilePath); err1 == nil {
			_, err = file.Write([]byte(template.Content))
			file.Close()
		} else {
			err = err1
			fmt.Println(err)
		}
	}
	return
}
