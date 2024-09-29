<picture>
  <p style="text-align:center;" align="center">
<a href="https://khulnasoft.com/meshplay">
<picture align="center">
<source media="(prefers-color-scheme: dark)" srcset="img\readme\meshplay-logo-dark-text-side.svg" width="70%" align="center" style="margin-bottom:20px;">
<source media="(prefers-color-scheme: light)" srcset="img\readme\meshplay-logo-light-text-side.svg" width="70%" align="center" style="margin-bottom:20px;">
<img alt="Shows an illustrated light mode meshplay logo in light color mode and a dark mode meshplay logo dark color mode." src="img\readme\meshplay-logo-light-text-side.svg" width="70%" align="center" style="margin-bottom:20px;"> </picture>
</a>

<br/><br/></p>
</picture>

# Meshplay Adapter for <adaptor-name>

[![Docker Pulls](https://img.shields.io/docker/pulls/khulnasoft/meshplay-<adaptor-name>.svg)](https://hub.docker.com/r/khulnasoft/meshplay-<adaptor-name>)
[![Go Report Card](https://goreportcard.com/badge/github.com/khulnasoft/meshplay-<adaptor-name>)](https://goreportcard.com/report/github.com/khulnasoft/meshplay-<adaptor-name>)
[![Build Status](https://github.com/khulnasoft/meshplay-<adaptor-name>/workflows/Meshplay-<adaptor-name>/badge.svg)](https://github.com/khulnasoft/meshplay-<adaptor-name>/actions)
[![GitHub](https://img.shields.io/github/license/khulnasoft/meshplay-<adaptor-name>.svg)](https://github.com/khulnasoft/meshplay-<adaptor-name>/blob/master/LICENSE)
[![GitHub issues by-label](https://img.shields.io/github/issues/khulnasoft/meshplay-<adaptor-name>/help%20wanted.svg)](https://github.com/khulnasoft/meshplay-<adaptor-name>/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22)
[![Website](https://img.shields.io/website/https/khulnasoft.com/meshplay.svg)](https://khulnasoft.com/meshplay)
[![Twitter Follow](https://img.shields.io/twitter/follow/khulnasoft.svg?label=Follow&style=social)](https://twitter.com/intent/follow?screen_name=khulnasoft)
[![Slack](https://img.shields.io/badge/Slack-@khulnasoft.svg?logo=slack)](http://slack.khulnasoft.com)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3564/badge)](https://bestpractices.coreinfrastructure.org/projects/3564)

<br />
<br />

<p style="clear:both;">
<h2><a href="https://khulnasoft.com/meshplay">Meshplay</a></h2>
<a href="https://meshplay.io"><img src="img/readme/meshplay-logo-light-text.svg"
style="margin:10px;" width="125px" 
alt="Meshplay - the Service Mesh Management Plane" align="left" /></a>
<a href="https://meshplay.io">Meshplay</a> is the multi-service mesh management plane offering lifecycle management of more types of service meshes than any other tool available today. Meshplay facilitates adopting, configuring, operating and managing performance of different service meshes and incorporates the collection and display of metrics from applications running on top of any service mesh. 
<br /><br /><p align="center"><i>If you’re using Meshplay or if you like the project, please <a href="https://github.com/khulnasoft/meshplay/stargazers">★</a> star this repository to show your support! 🤩</i></p>
</p>

<br />

<p style="clear:both;">
<h2><a name="contributing"></a><a name="community"></a> <a href="http://slack.khulnasoft.com">Community</a> and <a href="https://github.com/khulnasoft/khulnasoft/blob/master/CONTRIBUTING.md">Contributing</a></h2>
Our projects are community-built and welcome collaboration. 👍 Be sure to see the <a href="https://docs.google.com/document/d/17OPtDE_rdnPQxmk2Kauhm3GwXF1R5dZ3Cj8qZLKdo5E/edit">KhulnaSoft Community Welcome Guide</a> for a tour of resources available to you and jump into our <a href="http://slack.khulnasoft.com">Slack</a>! Contributors are expected to adhere to the <a href="https://github.com/cncf/foundation/blob/master/code-of-conduct.md">CNCF Code of Conduct</a>.

<a href="https://slack.meshplay.io">

<picture align="right">
  <source media="(prefers-color-scheme: dark)" srcset="img\readme\slack-dark-128.png"  width="110px" align="right" style="margin-left:10px;margin-top:10px;">
  <source media="(prefers-color-scheme: light)" srcset="img\readme\slack-128.png" width="110px" align="right" style="margin-left:10px;padding-top:5px;">
  <img alt="Shows an illustrated light mode meshplay logo in light color mode and a dark mode meshplay logo dark color mode." src="img\readme\slack-128.png" width="110px" align="right" style="margin-left:10px;padding-top:13px;">
</picture>
</a>

<a href="https://meshplay.io/community"><img alt="KhulnaSoft Cloud Native Community" src="img/readme/community.svg" style="margin-right:8px;padding-top:5px;" width="140px" align="left" /></a>

<p>
✔️ <em><strong>Join</strong></em> <a href="https://drive.google.com/open?id=1c07UO9dS7_tFD-ClCWHIrEzRnzUJoFQ10EzfJTpS7FY">weekly community meeting</a> on <a href="https://bit.ly/2SbrRhe">Fridays from 10am - 11am Central</a>.<br />
✔️ <em><strong>Watch</strong></em> community <a href="https://www.youtube.com/channel/UCFL1af7_wdnhHXL1InzaMvA?sub_confirmation=1">meeting recordings</a>.<br />
✔️ <em><strong>To Access Community Drive, </strong></em>fill the <a href="https://khulnasoft.com/newcomer">Community Member Form</a>.<br />
✔️ <em><strong>Discuss</strong></em> in the <a href="https://discuss.khulnasoft.com">Community Forum</a>.<br />
</p>
<p align="center">
<i>Not sure where to start?</i> Grab an open issue with the <a href="https://github.com/issues?q=is%3Aopen+is%3Aissue+archived%3Afalse+org%3Akhulnasoft+org%3Ameshplay+org%3Aservice-mesh-performance+org%3Aservice-mesh-patterns+label%3A%22help+wanted%22+">help-wanted label</a>.
</p>

<br />
<br />

## About KhulnaSoft

[KhulnaSoft](https://khulnasoft.com)'s cloud native application and infrastructure management software enables organizations to expect more from their infrastructure. We embrace developer-defined infrastructure. We empower engineer to change how they write applications, support operators in rethinking how they run modern infrastructure and enable product owners to regain full control over their product portfolio.

**License**

This repository and site are available as open source under the terms of the [Apache 2.0 License](https://opensource.org/licenses/Apache-2.0).
