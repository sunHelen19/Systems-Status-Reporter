<!-- PROJECT LOGO -->
<br />
<div align="center">
 
<h3 align="center">Systems Status Reporter</h3>

  <p text-align="center">
   Network multithreaded service for Statuspage
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
       </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

The application is a small network service that accepts requests over the network and returns data on the status of the company's systems. This data can be displayed on a web page and contain StatusPage, geography and service statuses.

<p text-align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

1. Clone the repo
   ```sh
   git clone https://gitlab.skillbox.ru/elena_fedorova_2/finalwork.git
      ```
   
2. Run the Simulator
   ```sh
   go run src/simulator/main.go
      ```
   
3. Run the application
   ```sh
   go run cmd/main.go
      ```

<p text-align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- USAGE EXAMPLES -->
## Usage

* to get JSON data, containing the result of the program with filtered and sorted data - enter the address in the browser:
  ```sh
  http://127.0.0.1:8282/json
  ```


* to display the received data on a web page (Status Page) - enter the address in the browser:
  ```sh
  http://127.0.0.1:8282/
  ```
    
<p text-align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p text-align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Project Link: [https://gitlab.skillbox.ru/elena_fedorova_2/finalwork](https://gitlab.skillbox.ru/elena_fedorova_2/finalwork)

<p text-align="right">(<a href="#readme-top">back to top</a>)</p>








