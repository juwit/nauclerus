# use golang for implementation

* status: accepted
* deciders: @juwit
* date: 2022-11-05

Technical Story: #1

## Context and Problem Statement

*nauclerus* aims to be a single-binary app, to be easy to deploy, update and operate.

*nauclerus* also needs to call Kubernetes APIs, in order to get information such as *Kubernetes* *namespace* list and *Helm*
*release* names and status.

## Considered Options

* develop in Java/Kotlin
* develop in golang

## Decision Outcome

Even if call *Kubernetes* APIs could be done in any language, developing in golang seems to be a better option.
As *Helm* SDK ships in golang, it would simplify development.