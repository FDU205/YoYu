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

type follows = {
    user_id:number, 
    username:string, 
}

type fans = {
    user_id:number, 
    username:string, 
}

export type {wallpost, messageBoxinfo, follows, fans};