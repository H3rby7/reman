# Install

!!! Follow [Installation Readme](INSTALLATION.md) !!!

## Database Setup

The first thing you need to do is `docker-compose up -d` and then 
open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

### Create Your Databases

(Creates the database inside the postgresql container)

Ok, so you've edited the "database.yml" file and started your database, now Buffalo can create the databases in that file for you:

	buffalo pop create -e development
	
### Update database schema

    buffalo pop migrate

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## Run finished image

Make sure to have your local database ready.
1. Run the postgres docker container as in docker-compose.yml
1. Run the buffalo commands as for local development, but use local-production environment (as flag: `-e local-production`)

`docker run -p 8080:3000 -e "DATABASE_URL=postgres://postgres:postgres@host.docker.internal:5432/reman_production?sslmode=disable" -e "SESSION_SECRET=joka" --user app:app test`

## What Next?

We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.

## DB Help

Start PSQL

    psql -U postgres -d database_name
    psql -U postgres -d reman_development
    
Switch database

    \c database_name
    
List relations:

    \dt
