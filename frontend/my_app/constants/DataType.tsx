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

type postinfo = {
    id:number,
    poster_id:number,
    message_box_id:number,
    content:string,
}

type follows = {
    user_id:number, 
    username:string, 
}

type fans = {
    user_id:number, 
    username:string, 
}

export type {wallpost, messageBoxinfo, postinfo, follows, fans};
