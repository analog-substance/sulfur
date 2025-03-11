---
title: sulfur
linkTitle: Docs
menu: {main: {weight: 20}}
---
> Attack Surface Management  
> [Docs](https://analog-substance.github.io/sulfur/) :: [Releases](https://github.com/analog-substance/sulfur/releases) :: [Code](https://github.com/analog-substance/sulfur/)  
> [![Go Tests](https://github.com/analog-substance/sulfur/actions/workflows/go-tests.yml/badge.svg)](https://github.com/analog-substance/sulfur/actions/workflows/go-tests.yml)

## Purpose

Maintain long term information about a target environment
* * *

## Features


## Installation

Download a [release from GitHub](https://github.com/analog-substance/sulfur/releases) or use Go Install:

```sh
go install github.com/analog-substance/sulfur@latest
```
* * *

## Feedback

### I have an issue or feature request

Sweet! [Open an issue](https://github.com/analog-substance/sulfur/issues/new) to start the conversation.

* * *

## Database layout (planned)


```mermaid
---
title: Attack Surface Management
---
classDiagram



    class Organization {
        -String name
        -String description
        +DateTime Created
        +DateTime Updated
    }
%%
%%    Organization -- Person
%%    class Person {
%%        -Reference Organization
%%        -String Name
%%        -String Description
%%        +DateTime Created
%%        +DateTime Updated
%%    }
%%
%%    Person -- PersonCredential
%%    class PersonCredential {
%%        -Reference Organization
%%        -Reference Person
%%        -String Username
%%        -String Password
%%        +DateTime Created
%%        +DateTime Updated
%%    }
%%
%%    Person -- PersonContactMethod
%%    class PersonContactMethod {
%%        -Reference Organization
%%        -Reference Person
%%        -String Name
%%        -String Value
%%        +DateTime Created
%%        +DateTime Updated
%%    }
%%
%%    Person -- PersonLinks
%%    class PersonLinks {
%%        -Reference Organization
%%        -Reference Person
%%        -String Name
%%        -String Link
%%        -String Type
%%        +DateTime Created
%%        +DateTime Updated
%%    }


    Organization -- Environment
    class Environment{
        -Reference Organization
        -String Name
        -String Description
        -DateTime Created
        -DateTime Updated
    }

    Environment -- EnvCIDR
    class EnvCIDR{
        -Reference Environment
        -String CIDR
        -int ASN
        -Bool is_private
        -Bool is_multicast
        -Bool is_unicast
        -Bool is_linklocal
        +DateTime LastSeen
        +DateTime Created
        +DateTime Updated
    }

    Environment -- EnvRootDomain
    class EnvRootDomain {
        -Reference environment
        -String Domain
        -String Registrar
        +DateTime LastSeen
        +DateTime Created
        +DateTime Updated
    }

%%    Organization -- Engagement
%%    class Engagement {
%%        -Reference Organization
%%        -String Name
%%        -String Description
%%        -String Status
%%        +DateTime ScheduledStart
%%        +DateTime ScheduledEnd
%%        +DateTime Created
%%        +DateTime Updated
%%        +Name()
%%        +Description()
%%    }
%%
%%    EngagementDoc -- Engagement
%%    class EngagementDoc {
%%        -Reference Engagement
%%        -string name
%%        -string type
%%    }
%%
%%    EngagementEnv -- Engagement
%%    EngagementEnv -- Environment
%%    class EngagementEnv {
%%        -Reference Organization
%%        -Reference Environment
%%    }
%%
%%    Lead -- Engagement
%%    class Lead {
%%        -Reference Engagement
%%        -String CVSS_Vector
%%        -String CVSS_Score
%%        -String CWE
%%        -string Title
%%        -string Summary
%%        +DateTime LastSeen
%%        +DateTime Created
%%        +DateTime Updated
%%    }
%%
%%    Lead -- LeadAsset
%%    LeadAsset -- IPService
%%    LeadAsset -- EnvRootDomain
%%    class LeadAsset {
%%        -Reference Engagement
%%        -Reference EnvRootDomain
%%        -Reference IPService
%%        -String URL
%%        +DateTime LastSeen
%%        +DateTime Created
%%        +DateTime Updated
%%    }

    EnvRootDomain -- DNSResolution
    class DNSResolution {
        -Reference EnvRootDomain
        +String Name
        +String Type
        +String Value
        +String TTL
        +DateTime LastResolved
        +string ResolveError
        +int ResolveErrorCount
        +DateTime LastSeen
        +DateTime Created
        +DateTime Updated
    }

    IPAddress -- DNSResolution
    IPAddress -- EnvCIDR
    class IPAddress {
        -String IPAddress
        -String Reverse
        -String Cloud
        -Bool is_cdn
        -Bool is_v6
        -Bool is_private
        -Bool is_multicast
        -Bool is_unicast
        -Bool is_linklocal
        +DateTime LastSeen
        +DateTime Created
        +DateTime Updated
    }

    class Certificate {
        -String fingerprint
        -string issuer
        -string valid_from
        -string expires_on
        -string subject
        -string alternative_subject
        
        +DateTime LastSeen
        +DateTime Created
        +DateTime Updated
    }

    Organization -- OrgCertificate
    OrgCertificate -- Certificate
    class OrgCertificate {
        -Reference Certificate
        -Reference Organization
    }

    IPService -- IPAddress
    class IPService {
        -Reference IPAddress
        -String Banner
        -String Service
        -string AppProto
        -string IPProtocol
        -int Port
        +DateTime LastSeen
        +DateTime Created
        +DateTime Updated
    }

    IPServiceCertificate -- IPService
    IPServiceCertificate -- Certificate
    class IPServiceCertificate {
        -Reference Certificate
        -Reference IPService
    }


    Whois -- EnvRootDomain
    class Whois {
        -Reference EnvRootDomain
        -String Org
        -String Service
        -string AppProto
        -string IPProtocol
        +DateTime LastSeen
        +DateTime Created
        +DateTime Updated
    }
```


