import {Movie} from "./movie";
import {Room} from "./room";

/**
 * Interface representing a showing
 */
export interface Showing {
    uid: string
    start_time: string
    end_time: string
    movie: Movie
    room: Room
}