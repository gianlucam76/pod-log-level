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