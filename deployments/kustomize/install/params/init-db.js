const mongoHost = process.env.AMBULANCE_API_MONGODB_HOST
const mongoPort = process.env.AMBULANCE_API_MONGODB_PORT

const mongoUser = process.env.AMBULANCE_API_MONGODB_USERNAME
const mongoPassword = process.env.AMBULANCE_API_MONGODB_PASSWORD

const database = process.env.AMBULANCE_API_MONGODB_DATABASE
const collection = process.env.AMBULANCE_API_MONGODB_COLLECTION

const retrySeconds = parseInt(process.env.RETRY_CONNECTION_SECONDS || "5") || 5;

// try to connect to mongoDB until it is not available
let connection;
while(true) {
    try {
        connection = Mongo(`mongodb://${mongoUser}:${mongoPassword}@${mongoHost}:${mongoPort}`);
        break;
    } catch (exception) {
        print(`Cannot connect to mongoDB: ${exception}`);
        print(`Will retry after ${retrySeconds} seconds`)
        sleep(retrySeconds * 1000);
    }
}

function initUserCollection() {
    collection = "users"
    
    const db = connection.getDB(database)
    db.createCollection(collection)

    // create indexes
    db[collection].createIndex({ "id": 1 })

    //insert sample data
    let result = db[collection].insertMany([{"id":"e6698483-7ef6-4432-acd9-baeb68830dae","name":"Natalie J. Capps","certifications":[{"id":"","userid":"","certificationid":"7df518c4-bff0-4026-8f52-6290a39d59c2","expiresat":"2022-03-03","issuedat":"2021-01-02"}]},{"id":"ab965b31-7f6d-42d4-9dde-366ae53d8d48","name":"Joan M. Thorp","certifications":[]},{"id":"c2f21653-98bc-489a-863e-e610688e8b00","name":"Victoria M. Armstrong","certifications":[]},{"id":"b8cf9a57-12ba-487f-a588-10d67d602709","name":"Miki M. Goff","certifications":[]},{"id":"5f1de7c2-eafc-4ac4-9ade-8384e88f29a4","name":"Dick K. Roberts","certifications":[]}]);

    if (result.writeError) {
        console.error(result)
        print(`Error when writing the data: ${result.errmsg}`)
    }
}

function initCertificationCollection() {
    collection = "certifications"
    
    const db = connection.getDB(database)
    db.createCollection(collection)

    // create indexes
    db[collection].createIndex({ "id": 1 })

    //insert sample data
    let result = db[collection].insertMany([{"id":"7df518c4-bff0-4026-8f52-6290a39d59c2","name":"Basic Life Support (BLS)","description":"Training in life-saving techniques including CPR, AED usage, and relief of choking for healthcare providers.","authority":"American Red Cross (ARC)"},{"id":"dcf01ae9-9b05-4276-aee7-2bd7cf75aea6","name":"Advanced Cardiovascular Life Support (ACLS)","description":"Advanced training in managing cardiac emergencies, including cardiac arrest, stroke, and acute coronary syndromes.","authority":"American Heart Association (AHA)"},{"id":"32703805-4a2c-4e5c-91f0-a4ba25b81dc4","name":"Pediatric Advanced Life Support (PALS)","description":"Training in managing pediatric emergencies, including cardiac arrest, respiratory distress, and shock.","authority":"American Heart Association (AHA)"},{"id":"d8d90583-3086-4fe0-b935-1a1445ba031b","name":"Emergency Medical Technician (EMT)","description":"Training in basic life support, patient assessment, and emergency care.","authority":"National Registry of Emergency Medical Technicians (NREMT)"},{"id":"1ddd60cd-befd-4927-b982-50301c7a059f","name":"Certified Medical Assistant (CMA)","description":"Certification for medical assistants demonstrating competency in clinical and administrative tasks in healthcare settings.","authority":"American Association of Medical Assistants (AAMA)"},{"id":"d4647038-792e-4033-9893-18cdd3cb1b58","name":"Registered Nurse (RN)","description":"Training in patient care, medication administration, and treatment planning.","authority":"American Nurses Credentialing Center (ANCC)"}]);

    if (result.writeError) {
        console.error(result)
        print(`Error when writing the data: ${result.errmsg}`)
    }
}

// if database and collection exists, exit with success - already initialized
const databases = connection.getDBNames()
if (databases.includes(database)) {
    const dbInstance = connection.getDB(database)
    collections = dbInstance.getCollectionNames()

    if (collections.includes("users")) {
      print(`Collection 'users' already exists in database '${database}'`)
    } else {
      initUserCollection();
    }

    if (collections.includes("certifications")) {
      print(`Collection 'certifications' already exists in database '${database}'`)    
    } else {
      initCertificationCollection();
    }
}

// exit with success
process.exit(0);
