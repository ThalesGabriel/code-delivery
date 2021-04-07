import { Prop, raw, Schema, SchemaFactory } from "@nestjs/mongoose";
import { Document } from "mongoose";

export type RouteDocument = Route & Document

@Schema()
export class Route {
    @Prop()
    _id: string;
    
    @Prop()
    title: string;

    @Prop(raw({
        lat: { raw: Number },
        lng: { raw: Number },
    }))
    startPosition: { lat: number, lng: number }

    @Prop(raw({
        lat: { raw: Number },
        lng: { raw: Number },
    }))
    endPosition: { lat: number, lng: number }
}

export const RouteSchema = SchemaFactory.createForClass(Route)