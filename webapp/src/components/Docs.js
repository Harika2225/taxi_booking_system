import React from 'react';

const DocsPage = () => {
  return (
    <div className="docs-page">
      <h1>Microservices Documentation</h1>
      <div className="docs-content">
        <p>Welcome to the Microservices Documentation. This guide provides insights into the microservices that power your application.</p>

        <h2>Understanding Your Microservices</h2>
        <p>
          Your application consists of multiple microservices collaborating to provide functionality. Currently, you have a total of 4
          microservices.
        </p>

        <h2>Getting Started in Development Mode</h2>

        <h4>Booting up the Required Services</h4>

        <p>
          If you have multiple microservices or a gateway, navigate to one of the gateway services. If your application architecture doesn't
          include a gateway service, navigate to any microservice.
        </p>
        <p>Once inside the gateway service or microservice, run the additional services necessary to launch your entire application.</p>

        <ol>
          <li>Open the terminal inside the navigated gateway/microservice.</li>
          <li>
            If applicable, start the database service: <code>npm run docker:db:up</code>
          </li>
          <li>
            If selected, start the OAuth service: <code>npm run docker:keycloak:up</code>
          </li>
          <li>
            If chosen, start the Eureka service: <code>npm run docker:jhipster-registry:up</code>
          </li>
        </ol>
        <p>Note: If you have both a gateway and microservices, navigate to the gateway service to run the above commands.</p>

        <h4>Booting up the UI Services</h4>
        <ol>
          <li>Navigate to the UI application.</li>
          <li>Open the terminal inside the navigated UI service.</li>
          <li>
            Run: <code>npm start</code>
          </li>
        </ol>

        <h4>Booting up the Gateway/Microservices (Spring Boot)</h4>
        <ol>
          <li>Navigate to each gateway/microservice individually.</li>
          <li>Open the terminal inside the navigated service.</li>
          <li>
            Run: <code>./mvnw</code>
          </li>
        </ol>

        <h2>Resolving Cross-Origin Issues in Development Mode</h2>
        <p>
          If you encounter cross-origin issues while running the UI application in the development environment, don't worry. Here's how to
          address it:
        </p>

        <ol>
          <li>Open your Spring Boot application.</li>
          <li>
            Navigate to: <code>src/main/resources/config</code>.
          </li>
          <li>
            Locate the <code>application-dev.yml</code> file.
          </li>
          <li>
            Add your UI application's origin to the <code>allowed-origins</code> section:
          </li>
        </ol>

        <pre>
          cors:
          <br />
          &nbsp;&nbsp;allowed-origins: 'http://your-ui-origin'
        </pre>

        <p>
          Remember to replace <code>http://your-ui-origin</code> with your UI application's actual URL.
        </p>

        <pre>
          Just for your reference on how to add allowed-origins:
          <br />
          <br />
          cors:
          <br />
          &nbsp;&nbsp;allowed-origins: 'http://localhost:3000,http://localhost:4200'
        </pre>

        <p>By following these steps, you'll ensure a smooth cross-origin experience during development.</p>
      </div>
    </div>
  );
};

export default DocsPage;
