type wallpost = {
    id:number, 
    poster_id:number, 
    poster_name:string, 
    content:string,
    visibility:number,
}

type messageBoxinfo = {
    id:number, 
    owner_id:number,
    title:string,
    owner_name:string,
}

export type {wallpost, messageBoxinfo};