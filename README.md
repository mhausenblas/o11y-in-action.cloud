## Cloud Observability in Action

Welcome to Cloud Observability in Action, your hands-on guide to applying
observability in the context of cloud native environments.

> Observability is the capability to continuously generate and discover 
> actionable insights based on signals from the system under observation 
> with the goal to influence the system.

![MEAP cover](co11yia-meap-cover.png)

In this book you will learn about the basic signal types (logs, metrics, traces,
profiles), telemetry including agents, back-end and front-end destinations, 
and goood practices around dashboarding, alerting, and SLOs/SLIs.

Some chapters of the book are now available via the [Manning MEAP Program](https://www.manning.com/books/cloud-observability-in-action)
and you can find [code snippets](https://github.com/mhausenblas/o11y-in-action.cloud/tree/main/code) we use throughout the book via the 
site you're on, currently.

The WIP table of contents looks as follows:

## Chapter 1: End-to-end Observability Example
In the context of this book we focus on cloud native environments such as 
Kubernetes and serverless offerings (such as FaaS like AWS Lambda). We mainly
use open source observability tooling (Grafana, Prometheus, Jaeger) so that 
you can try out everything without license costs. While it is important that
we use open source tooling to show the concepts in action, they are universally
applicable (that is, using any of the commerical offerings). 
In this chapter we have a look at an end-to-end example and define terminology,
from sources to agents to destinations.

1. What is Observability?
1. Roles and Goals
1. Example Microservices App
1. Challenges and How Observability Helps

## Chapter 2: Signal Types
In this chapter we review different signal types most often used, 
how to instrument and collect each, and discuss the costs and benefits of doing 
that. With observability you want to take an Return-On-Investment (ROI) driven
approach. In other words, you need to understand the costs of each signal type 
and what it enables you to do.

1. Reference Example
1. Assessing Instrumentation Costs
1. Logs
1. Metrics
1. Traces
1. Selecting Signals

## Chapter 3: Sources
This chapter covers signal sources. We discuss the type of sources that exist
and when to select which source, how you can gain actionable insights from selecting
the right sources for a task and how to deal with code you own including supply
chain aspects.

1. Selecting Sources
1. Compute-related Sources
1. Storage-related Sources
1. Network-related Sources
1. Your Code

## Chapter 4: Agents
In this chapter we discuss instrumentation and review different agents,
from log routers to OpenTelemetry. You will learn how to select and use agents
with an emphasis on what OpenTelemetry brings to the table for unified telemetry,
including correlation of signals.

1. Log Routers
1. Metrics Collection
1. OpenTelemetry
1. Other Agents
1. Selecting An Agent

## Chapter 5: Back-end Destinations
ETA: 08/2022

## Chapter 6: Front-end Destinations
ETA: 09/2022

## Chapter 7: Alerting
ETA: 09/2022

## Chapter 8: Distributed Tracing
ETA: 10/2022

## Chapter 9: Continuous Profiling
ETA: 11/2022

## Chapter 10: Service Level Objectives
ETA: 11/2022

