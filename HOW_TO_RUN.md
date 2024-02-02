# HOW TO RUN

Your prototype is composed of the following components:

### UI (Client frameworks)

-   [webapp](webapp/README.md)

### Backend (Microservices)

-   [bookingmanagementmodule](bookingmanagementmodule/README.md)
-   [customermanagementmodule](customermanagementmodule/README.md)
-   [drivermanagementmodule](drivermanagementmodule/README.md)
-   [paymentmanagementmodule](paymentmanagementmodule/README.md)

### Infrastructure as a Code

If a Deployment Infrastructure is chosen while generating code, additional 2 components will be generated to deploy the prototype.

-   Terraform (Contains IaaC to create and manage resources)
-   Kubernetes (Contains manifest files for the you prototype components)
-   _auto-deploy_ & _auto-destroy_ scripts in Terraform directory (Helper scripts for rapid orchestration)

### Common Prerequisites

-   Docker Engine >= 24.0
-   Docker Desktop (optional) >= 4.25

### Please boot up the components in the below order to run the prototype:

-   bookingmanagementmodule
-   customermanagementmodule
-   drivermanagementmodule
-   paymentmanagementmodule
-   webapp

Please refer to the component level README's for component level details.
