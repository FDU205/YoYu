const rootUrl = 'http://192.168.1.109:8080/api';

async function getData(url: string, token = ""): Promise<any> {
  url = rootUrl + url;
  console.log("token: "+token);
  const myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");
  if(token !== "") {
    myHeaders.append("Authorization", "BEARER " + token);
  }

  const response = await fetch(url, {
    method: 'GET',
    headers: myHeaders
  });
  return response.json();
}

async function postData(url: string, body: any, token = ""): Promise<any> {
  url = rootUrl + url;
  console.log("token: "+token);
  const myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");
  if(token !== "") {
    myHeaders.append("Authorization", "BEARER " + token);
  }

  const response = await fetch(url, {
    method: 'POST',
    headers: myHeaders,
    body: JSON.stringify(body)
  });
  return response.json();
}

async function putData(url: string, body: any, token = ""): Promise<any> {
  url = rootUrl + url;

  const myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");
  if(token !== "") {
    myHeaders.append("Authorization", "BEARER " + token);
  }

  const response = await fetch(url, {
    method: 'PUT',
    headers: myHeaders,
    body: JSON.stringify(body)
  });
  return response.json();
}

async function deleteData(url: string, token = ""): Promise<any> {
  url = rootUrl + url;

  const myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");
  if(token !== "") {
    myHeaders.append("Authorization", "BEARER " + token);
  }

  const response = await fetch(url, {
    method: 'DELETE',
    headers: myHeaders
  });
  return response.json();
}

export { getData, postData, putData, deleteData };