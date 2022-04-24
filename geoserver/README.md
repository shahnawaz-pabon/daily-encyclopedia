<div align="center">
  <a href="http://geoserver.org/">
    <img alt="geoserver" src="../logos/geoserver.png" height="96" width="96"/>
  </a>
  <h1>GeoServer</h1>
</div>

# Table of Contents

- [Setup GEO Server on Ubuntu](#setup-geo-server-on-ubuntu)
- [Setup POSTGIS with Docker](#setup-postgis-with-docker)

## Setup GEO Server on Ubuntu

Make sure you have a `Java Runtime Environment (JRE)` installed on your system. GeoServer requires a Java 8 or Java 11 environment, available from [OpenJDK](https://openjdk.java.net/), [Adoptium](https://adoptium.net/), or provided by your OS distribution.
Navigate to the [GeoServer Download page](http://geoserver.org/download).
Select the version of GeoServer that you wish to download. If you’re not sure, select [Stable](http://geoserver.org/release/stable) release.
Select **`Platform Independent Binary`** on the download page: [geoserver-2.20-SNAPSHOT-bin.zip](https://build.geoserver.org/geoserver/2.20.x/geoserver-2.20.x-latest-bin.zip)
Download the zip archive and unpack to the directory where you would like the program to be located.

Note: A suggested location would be `/usr/share/geoserver`.

Add an environment variable to save the location of GeoServer by typing the following command:

```sh
echo "export GEOSERVER_HOME=/usr/share/geoserver" >> ~/.bashrc
. ~/.bashrc
```

Make yourself the owner of the `geoserver` folder. Type the following command in the terminal window, replacing `USER_NAME` with your own username:

```sh
sudo chown -R USER_NAME /usr/share/geoserver/
```

Start GeoServer by changing into the directory `geoserver/bin` and executing the `startup.sh` script:

```sh
cd geoserver/bin
sh startup.sh
```

In a web browser, navigate to [http://localhost:8080/geoserver](http://localhost:8080/geoserver).
The default `username: admin` and `password: geoserver`
To shut down GeoServer, either close the persistent command-line window or run the `shutdown.sh` file inside the `bin` directory.

## Setup POSTGIS with Docker

```sh
docker run -d -p 5432:5432 --name db kartoza/postgis:13.0
```

Postgis docker’s Credentials

**Username: `docker`**
**Password: `docker`**
**Database: `gis`**
