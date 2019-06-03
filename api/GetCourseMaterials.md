**Get Course Materials**
----
    Returns the list of teaching materials for a given course
* **URL**

  /teachingMaterials/:courseId
  
* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `courseId=[string]`
   
   * **Data Params**

    None

* **Success Response:**

  * **Code:** 200 OK <br />
    **Content:** 
    `[
    "5cecf333b6254b20162a9d24_file1.txt",
    "5cecf333b6254b20162a9d24_file2.txt"
    ]` This is the list of available contents for the given course
 
* **Error Response:**

  * **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `{ error : "Api Gateway - Internal Server Error" }`
  
  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "Not Held Course"}`
    This is returned when the teacher doesn't hold the course with
    the provided identifier
    
  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "No token provided" }`
    
  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "Wrong credentials" }`
    
  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "Expired token" }`