**Generate temporary download link**
----
    Generate a link useful to download the file with specified filename from the data-store.
    The link has temporary validity.
* **URL**

  /downloadLink/:filename/course/:courseId
  
* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `filename=[string]`
   
   `filename` is the name of the file to upload
   
   `courseId=[string]`
   
   `courseId` is the course envolved in the upload


* **Data Params**

    None

* **Success Response:**

  * **Code:** 200 OK <br />
    **Content:** `[
    https://teachingmaterialmanagement.s3.amazonaws.com/test.txt?X-Amz-Security-Token=FQoGZXIvYXdzEBEaDNHo%2BoMaGkiibuk%2BlCKNA6lPFfGSGoxQkOLcgO3TH%2F84Dl%2FUd1pfHyTaYVTfBFra1ISGQXnqy9iCdQH2Q5S7p%2FkU7KLqmOsVW6kUBp1%2BQGmZCv19S6OwrwCLXw06aycPSBcgYkIu%2BxyHxBKa%2Fius%2B70PwDcvMIcRumiva7cd1A%2FRKMNHBuSdA0BkVffCRPBWprBVNP4ljVLMOVg5FVudgcskFT19o8YBK3X0TEMwqZV4JvHicDrgQ6E9ITRwC5QB7PnbvhSFxiMgQVaGajg0%2FIdQyTytGRssRWmXWC9SMUBffCSPP4F811a0RFTtnVif%2BdZdsNDSWgLCkMk7M8niOBzQCm727lD1zHzJQhmkDAInO708uyivt%2FJNn%2BG7bOZut%2BvDXY%2FJt6naydD6%2FQbeBFib8Bn8X9G0mVA3b7gqACIu0fHdA7JcceAagUFlJZrWGQi6tiwidKQ0ZfgoBW%2BCpoTvipqS69kRwuNYPUxoqffqvhiU11ssKgN%2B1TeRlPZjy67A3yHHl66f0HvZTQ5DjF0hBXIaReTcva1%2BXxso6eqq5wU%3D&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20190526T162143Z&X-Amz-SignedHeaders=host&X-Amz-Expires=599&X-Amz-Credential=ASIASR7ANGZQ4HXNPKNI%2F20190526%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Signature=e28e38672d6404c14eecd5ae2cd9898d033878297d5dec5fffdbf3b9232e7e3b
    ]`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "Not found"}`
    
  OR
  
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
