const rootUrl = 'http://192.168.1.105:4523/m1/2465803-0-default/api';

async function getData(url: string, token = ""): Promise<any> {
  url = rootUrl + url;

  const myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");
  if(token !== "") {
    myHeaders.append("Authorization", "Bearer " + token);
  }

  const response = await fetch(url, {
    method: 'GET',
    headers: myHeaders
  });
  return response.json();
}

async function postData(url: string, body: any, token = ""): Promise<any> {
  url = rootUrl + url;

  const myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");
  if(token !== "") {
    myHeaders.append("Authorization", "Bearer " + token);
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
    myHeaders.append("Authorization", "Bearer " + token);
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
    myHeaders.append("Authorization", "Bearer " + token);
  }

  const response = await fetch(url, {
    method: 'DELETE',
    headers: myHeaders
  });
  return response.json();
}

export { getData, postData, putData, deleteData };