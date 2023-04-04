# pod-log-level
A small library to allow changing pod log level without restarting the pod.

Sometimes it is important to be able to change Pod log level at run-time, without restarting the Pod. Here is where this library will help. 

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
		"projectsveltos", "SveltosManager", ctrl.Log.WithName("log-setter"),
		ctrl.GetConfigOrDie())
```

To change log level, you have two choices. You can manually edit "default" LogSetting instance or better, use CLI that comes with this library.

```bash
git clone git@github.com:gianlucam76/pod-log-level.git
```

```bash
make build
```

```bash
./bin/manager log-level set --namespace=projectsveltos --identifier=SveltosManager --info
```

When you do that, you can see log level is changed at runtime. Snippet from pod logs

```
I0404 13:41:23.792082       1 log_settings.go:177] "log-setter: got add notification for LogSettings"
I0404 13:41:23.793416       1 log_settings.go:247] "log-setter: Setting log severity to info" default="0"
```

You can see all settings

```bash
./bin/manager log-level show                                                             
+---------------------+----------------------+--------------+
| COMPONENT NAMESPACE | COMPONENT IDENTIFIER |  VERBOSITY   |
+---------------------+----------------------+--------------+
| projectsveltos      | projectsveltos       | LogLevelInfo |
+---------------------+----------------------+--------------+
```

You can increase log level to debug for instance

```bash
 ./bin/manager log-level set --namespace=projectsveltos --identifier=SveltosManager --verbose
 ```

```bash
 ./bin/manager log-level show                                                                
+---------------------+----------------------+-----------------+
| COMPONENT NAMESPACE | COMPONENT IDENTIFIER |    VERBOSITY    |
+---------------------+----------------------+-----------------+
| projectsveltos      | projectsveltos       | LogLevelVerbose |
+---------------------+----------------------+-----------------+
```

