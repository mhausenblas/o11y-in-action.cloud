## Cloud Native Observability in Action

Welcome to Cloud Native Observability in Action, your hands-on guide to applying
observability in the context of cloud native environments.

> Observability is the capability to continuously generate and discover 
> actionable insights based on signals from the system under observation 
> with the goal to influence the system.

![MEAP cover](co11yia-meap-cover.png)

In this book you will learn about the basic signal types (logs, metrics, traces,
profiles), telemetry including agents, back-end and front-end destinations, 
and goood practices around dashboarding, alerting, and SLOs/SLIs.

The first five chapters of the book are now available via the
[Manning MEAP Program](https://www.manning.com/books/cloud-observability-in-action)
and you can find all the [code snippets](https://github.com/mhausenblas/o11y-in-action.cloud/tree/main/code) 
we use for the hands-on exercises throughout the book via the 
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
The chapter focuses on back-ends as the source of truth for your telemetry
signals. You will learn to use and select back-ends for logs, metrics, and
traces with deep dives on TSDBs and column-oriented datastores such as
ClickHouse.

1. Back-end Destinations Terminology
1. Back-end Destinations for Logs
1. Back-end Destinations for Metrics
1. Back-end Destinations for Traces
1. Columnar Datastores
1. Selecting Back-End Destinations

## Chapter 6: Front-end Destinations
In the chapter we talk about front-ends as the place where you consume
the telemetry signals. You will learn about pure front-ends and all-in-ones
and how to go about selecting them.

1. Front-ends (Grafana, Kibana/OpenSearch Dashboards)
1. All-in-ones (Jaeger, Pixie, commerical offerings)
1. Selecting Front-ends and All-in-ones

## Chapter 7: Cloud Operations
This chapter covers an aspect of cloud native solutions called "cloud operations" 
including how to detect when something is not working the way that it should, 
how to react to abnormal behavior, and how to learn from previous mistakes. 
You will also learn about alerting, usage, and cost tracking.

1. How to manage incidents
1. Health monitoring and alerts
1. Governance, usage, and cost tracking

## Chapter 8: Distributed Tracing
ETA: 02/2023

## Chapter 9: Developer Observability
ETA: 03/2023

## Chapter 10: Service Level Objectives
ETA: 03/2023

