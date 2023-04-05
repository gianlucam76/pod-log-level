# pod-log-level
A small library to allow changing pod log level without restarting the pod.

Sometimes it is important to be able to change Pod log level at run-time, without restarting the Pod. Here is where this library will help. 

Currently works with k8s.io/klog/v2 

## How to use it
Deploy LogSetting CRD

```
kubectl apply -f https://raw.githubusercontent.com/gianlucam76/pod-log-level/main/config/crd/bases/open.projectsveltos.io_logsettings.yaml
```

Then in your application import library and register:

```go

        import "github.com/gianlucam76/pod-log-level/lib"
...
	lib.RegisterForLogSettings(ctx,
		"<YOUR POD NAMESPACE>", "<YOUR POD IDENTIFIER>", ctrl.Log.WithName("log-setter"),
		ctrl.GetConfigOrDie())
```

Make sure pod identifier does not change across pod restarts.

Make sure ServiceAccount associated to your Pod has permission to get/list/watch LogSettings

```
- apiGroups:
  - open.projectsveltos.io
  resources:
  - logsettings
  verbs:
  - get
  - list
  - watch
```

That's all that is required.

## Example

```
	lib.RegisterForLogSettings(ctx,
		"projectsveltos", "SveltosManager", <logr.Logger>,
		<cluster *rest.Config>)
```

To change log level, you have two choices. You can manually edit "default" LogSetting instance or better, use CLI that comes with this library.

```bash
git clone git@github.com:gianlucam76/pod-log-level.git
```

```bash
make build
```

```bash
./bin/helper log-level set --namespace=projectsveltos --identifier=SveltosManager --info
```

When you do that, you can see log level is changed at runtime. Snippet from pod logs

```
I0404 13:41:23.792082       1 log_settings.go:177] "log-setter: got add notification for LogSettings"
I0404 13:41:23.793416       1 log_settings.go:247] "log-setter: Setting log severity to info" default="0"
```

You can see all settings

```bash
./bin/helper log-level show                                                             
+---------------------+----------------------+--------------+
| COMPONENT NAMESPACE | COMPONENT IDENTIFIER |  VERBOSITY   |
+---------------------+----------------------+--------------+
| projectsveltos      | SveltosManager       | LogLevelInfo |
+---------------------+----------------------+--------------+
```

You can increase log level to debug for instance

```bash
 ./bin/helper log-level set --namespace=projectsveltos --identifier=SveltosManager --verbose
 ```

```bash
 ./bin/manager log-level show                                                                
+---------------------+----------------------+-----------------+
| COMPONENT NAMESPACE | COMPONENT IDENTIFIER |    VERBOSITY    |
+---------------------+----------------------+-----------------+
| projectsveltos      | SveltosManager       | LogLevelVerbose |
+---------------------+----------------------+-----------------+
```

## Log levels

By default log levels are:

1. info => V(0)
2. debug => V(5)
3. verbose => V(10)

Library exposes API to change those setting per Pod.

```go
	instance := lib.GetInstance()
	instance.SetInfoValue(2)
	instance.SetDebugValue(6)
	instance.SetVerboseValue(8)
```

then  if you set level to Debug

```bash
./bin/helper log-level set --namespace=projectsveltos --identifier=SveltosManager --debug  
```

in your POD logs you can see level is 6

```bash
I0404 15:14:16.293690       1 log_settings.go:198] "log-setter: got update notification for LogSettings"
I0404 15:14:16.293864       1 log_settings.go:232] "log-setter: Setting log severity to debug" debug="6"
```
