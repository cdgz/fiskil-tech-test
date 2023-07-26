# Part 1

> As an engineer, what steps would you take to address this challenge? Please walk me through your approach.

We are talking about multitude of single-tenant SaaS instances on GCP.
My approach would be creating microservice architecture hosted on GKE, with services exposed to WAN through ingresses and associated load balancer(s) and dedicated DNS entry(-ies). For scalability (and potential privacy) reasons, we might need to encode customer names into short hashes, prepended to top-level DNS domain.

Assuming that one single SaaS instance would involve multiple services communicating with each other - as
well as some GCP-managed services like SQL/object storage/AI - I'do go with one single GKE plane per customer.
Since we are in B2B context, it would allow better isolation and segregation, ensuring clearer perimeter and more explicit blast-radius for each customer environment which is great from security perspective.
Also this approach facilitates the rollout of SaaS solution in any geographical location where GCP is present.

We might need to define "control instance" which would act as a global/main installation having knowledge of other satellite instances, and some sort of API endpoint which would allow to register / activate new regions.

> Furthermore, please provide a high-level component diagram using cloud-based design principles, utilising components from GCP and open source tools that would support your proposed solution.

![Screenshot from 2023-07-26 11-51-47](https://github.com/cdgz/fiskil-tech-test/assets/734701/47de72d0-02e1-4027-b228-4ab5ee0c7164)

![Screenshot from 2023-07-26 11-55-06](https://github.com/cdgz/fiskil-tech-test/assets/734701/f7fa2662-a82f-4929-a587-14157cd67a11)

On simplified diagrams, we run every SaaS instance in dedicated GKE plane. Compute worfkload hosts both UI and backend services, which are exposed as K8s services to WAN via ingresses tied to Google Cloud Load Balancers who terminate TLS. Compute workload has access to any storage resources (Google Cloud SQL, Cloud Storage buckets, Datastore/Bigtable) and any other managed services like Cloud Run, Cloud Functions, Pub-Sub, AI.

Every customer would access DNS endpoint only known to them. As additinal security, WAF rules might be installed to allow incoming traffic only from IP addresses associated to given customer.

Scaling up and out for each customer would be done within SaaS instance perimeter, with help of autoscaling of compute workload and horizonal growth of storage layers (e.g. adding more nodes).

> Finally, how would you approach continuous integration and continuous deployment (CI/CD) for this SaaS-based application? What advice would you offer the company to ensure successful implementation and maintenance of the application?

Code-wise, for each service I would implement workflow with 2 persistent branches (development and main) with regular merges of development into main before release.

At CI level: 
- We'd test every PR with a lightweight set of unit tests scoped to given service.
- Upon merge of PR, a more extended (integration) test suite would kick in, where interaction with other services is checked.
- Upon release on either of branches, a deliverable (typically a package/docker image) would be pushed to artifact storage. Security scans would block the process if any critical issues are found at this point.

At CD level, we'd operate with above deliverables and promote those packages between environments. This is the stage where all parameter customization happens (environment name, customer ID, compute workload types) and secrets are injected into runtime environment. An additional post-deploy end-to-end healthcheck for the whole stack is very useful here.
