import {File} from "./file";
import {Genre} from "./genre";
import {Classification} from "./classification";

/**
 * Interface representing a movie
 */
export interface Movie {
    uid: string
    name: string
    description: string
    total_time: number // in minutes
    release_date: string
    status: string

    files: File[]
    genres: Genre[]
    classifications: Classification[]
}