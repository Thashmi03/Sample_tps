import http from 'k6/http';
import { check, sleep } from 'k6';
export const options = {
  vus:1,
  iterations :100000
}
export default function () {
  // // http.post('http://localhost:4000/test');
  // // Define the API endpoint
  // // const url = 'http://localhost:4000/test';
  const url = 'http://localhost:4000/test'
  // // Define the JSON request body
  const payload = JSON.stringify({
    "name": "thash",
    // "id":18
  });
  // const jsonData = '{"name":"thash","id":318}';

  // Make an HTTP POST request with the JSON payload

  const headers= {
    'Content-Type': 'application/json',
    Authorization: 'Bearer eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJVc2VybmFtZSI6IkphdmFJblVzZSJ9.mdVt7B4035pO0xLJtw9x47MPaYlt5YNyNl6nHAY2Uvo',
  }
 


  // Define HTTP headers
//   const headers = {
//     'Content-Type': 'application/json',
//     Authorization: 'Bearer eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJVc2VybmFtZSI6IkphdmFJblVzZSJ9.mdVt7B4035pO0xLJtw9x47MPaYlt5YNyNl6nHAY2Uvo', // Replace with a valid token
// };
  // Send a POST request to the API
  const res = http.post(url, payload, { headers });

  // Check if the response status code is 200
  check(res, {
    'Status is 200': (r) => r.status === 200,
  });

}

