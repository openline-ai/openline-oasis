<div align="center">
  <a href="https://openline.ai">
    <img
      src="https://www.openline.ai/TeamHero.svg"
      alt="Openline Logo"
      height="64"
    />
  </a>
  <br />
  <p>
    <h3>
      <b>
        Openline Oasis app
      </b>
    </h3>
  </p>
  <p>
    Openline customerOS is the easiest way to consolidate, warehouse, and build applications with your customer data.
  </p>
  <p>

[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen?logo=github)][oasis-repo] 
[![license](https://img.shields.io/badge/license-Apache%202-blue)][apache2] 
[![stars](https://img.shields.io/github/stars/openline-ai/openline-oasis?style=social)][oasis-repo] 
[![twitter](https://img.shields.io/twitter/follow/openlineAI?style=social)][twitter] 
[![slack](https://img.shields.io/badge/slack-community-blueviolet.svg?logo=slack)][slack]

  </p>
  <p>
    <sub>
      Built with ❤︎ by the
      <a href="https://openline.ai">
        Openline
      </a>
      community!
    </sub>
  </p>
</div>


## 👋 Overview

TBD

## 🚀 Installation

### set up a minikube envoronment

* This process has been tested in both macos and Ubuntu 20.04 the install process may need to be adapted for other platforms
* If you use codespaces, be sure to use the 4 core environment
* The voice network can not run on arm64, if arm64 is detected kamailio and asterisk will not be installed


#### Set up your mail configuraiton
Open the file
```
deployment/k8s/local-minikube/apps-config/channels-api.yaml
```
Update the 
* SMTP_SERVER_ADDRESS
* SMTP_SERVER_USER
* SMTP_SERVER_PASSWORD
* SMTP_FROM_USER


To setup the environment for inbound email you must follow the procedure [here](channels-api/ses-forwarder-lambda/README.md) to set up Amazon SES to send e-mails to you

Once these steps are completed you can proceed to the deployment step

#### Run the deployment script

```
deployment/k8s/local-minikube/0-build-deploy-openline-oasis-local-images.sh build
```

after the script completes you can validate the status of the setup by running
```
kubectl -n openline-development get pod
```

if you are not running the minikube on your local machine and the minikube is behind a nat you will probably need to install the turn server to have audio
```
./1-start-turn.sh
```
To start the tunnels for the core services you can run the folowing
```
./2-start-tunnels.sh
```

## 🙌 Features

TBD

## 🤝 Resources

- For help, feature requests, or chat with fellow Openline enthusiasts, check out our [slack community][slack]!
- Our [docs site][docs] has references for developer functionality, including the Graph API

## 💪 Contributions

- We love contributions big or small!  Please check out our [guide on how to get started][contributions].
- Not sure where to start?  [Book a free, no-pressure, no-commitment call][call] with the team to discuss the best way to get involved.

## ✨ Contributors

A massive thank you goes out to all these wonderful people ([emoji key][emoji]):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center"><a href="https://github.com/tsearle"><img src="https://avatars.githubusercontent.com/u/4540323?v=4?s=100" width="100px;" alt="tsearle"/><br /><sub><b>tsearle</b></sub></a><br /><a href="https://github.com/openline-ai/openline-oasis/commits?author=tsearle" title="Code">💻</a> <a href="https://github.com/openline-ai/openline-oasis/commits?author=tsearle" title="Documentation">📖</a></td>
      <td align="center"><a href="https://github.com/gabriel0003"><img src="https://avatars.githubusercontent.com/u/108665351?v=4?s=100" width="100px;" alt="Gabriel Gontariu"/><br /><sub><b>Gabriel Gontariu</b></sub></a><br /><a href="https://github.com/openline-ai/openline-oasis/commits?author=gabriel0003" title="Code">💻</a> <a href="https://github.com/openline-ai/openline-oasis/pulls?q=is%3Apr+reviewed-by%3Agabriel0003" title="Reviewed Pull Requests">👀</a></td>
      <td align="center"><a href="https://github.com/edifirut"><img src="https://avatars.githubusercontent.com/u/108661145?v=4?s=100" width="100px;" alt="edifirut"/><br /><sub><b>edifirut</b></sub></a><br /><a href="https://github.com/openline-ai/openline-oasis/pulls?q=is%3Apr+reviewed-by%3Aedifirut" title="Reviewed Pull Requests">👀</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

## 🪪 License

- This repo is licensed under [Apache 2.0][apache2], with the exception of the ee directory (if applicable).
- Premium features (contained in the ee directory) require an Openline Enterprise license.  See our [pricing page][pricing] for more details.


[apache2]: https://www.apache.org/licenses/LICENSE-2.0
[call]: https://meetings-eu1.hubspot.com/matt2/customer-demos
[careers]: https://openline.ai
[contributions]: https://github.com/openline-ai/community/blob/main/README.md
[docs]: https://openline.ai
[emoji]: https://allcontributors.org/docs/en/emoji-key
[oasis-repo]: https://github.com/openline-ai/openline-customer-os/
[pricing]: https://openline.ai/pricing
[slack]: https://join.slack.com/t/openline-ai/shared_invite/zt-1i6umaw6c-aaap4VwvGHeoJ1zz~ngCKQ
[twitter]: https://twitter.com/OpenlineAI
