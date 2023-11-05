import { type User } from "@/services/user";
export interface Video {
    video_id: bigint,
    author:User,
    play_url:string,
    cover_url:string,
    favorite_count:number,
    collection_count:number,
    comment_count:number,
    is_favorite:boolean,
    title:string,
    partition:number,
    create_time:Date,
    // videoURL?: string,
    // videoCoverURL?: string,
    // videoTile?: string,
    // username?: string,
    // publishTime?: number,
    // likeCount?: number,
}