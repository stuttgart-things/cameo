# stuttgart-things/cameo

## Working with Cameo

Cameo is a series of functions to manage the [Cameo-API](https://codehub.sva.de/Lab/stuttgart-things/dev/cameo-api). With these functions, you are able to add, delete, update or list the values within a redis account, under a specified redis key. 

In order for the functions to work, the "Cameo-API" should be running and the connection data (Host and port) should be known.

<details><summary><b>cameo list</b></summary>


Run the following command to list the values within the set. You can either list all values or a few random values.

```
HOST=<HOST> PORT=<PORT> cameo list
```

Example for substituted values:

```
HOST=http://localhost PORT=5001 cameo list
```
</details>

<details><summary><b>cameo add</b></summary>


Run the following command to add new values to the set.

```
HOST=<HOST> PORT=<PORT> cameo add
```

Example for substituted values:

```
HOST=http://localhost PORT=5001 cameo add
```
</details>

<details><summary><b>cameo update</b></summary>


Run the following command to update an old value and change it to a new values.

```
HOST=<HOST> PORT=<PORT> cameo update
```

Example for substituted values:

```
HOST=http://localhost PORT=5001 cameo update
```
</details>

<details><summary><b>cameo delete</b></summary>


Run the following command to delete a values from the set.

```
HOST=<HOST> PORT=<PORT> cameo delete
```

Example for substituted values:

```
HOST=http://localhost PORT=5001 cameo delete
```
</details>


## Redis Account - Cameo-API connection

The credentials for the redis account into which you will connect are defined within the enviornment variables of "Cameo-API". There you can find and modify the following variables:
- **REDIS_KEY**  - The name of the key where the desired values are stored.
- **REDIS_PORT**  - The port through which the connection with redis will be established.
- **REDIS_SERVER**  - The host with which the connection is desired.
- **REDIS_PASSWORD**  - The password to connect with your redis account.

Additionally, you can also modify the port through which you communicate with "Cameo-API"
- **API_PORT_LOCAL** 

Author Information
------------------
Calva, Ana
(SVA GmbH, 03/2023)

