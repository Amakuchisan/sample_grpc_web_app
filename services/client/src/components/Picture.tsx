import { useState } from 'react';
import { GetPicturesRequest, GetPicturesReply, StreamGetPicturesReply } from "../pb/picture/picture_pb";
import { PictureClient } from "../pb/picture/PictureServiceClientPb";
import { ClientReadableStream, Error } from 'grpc-web';

export const Picture = () => {
    const [num, setNumber] = useState(1); // 枚数の指定
    const [pictures, setPictures] = useState<JSX.Element[]>([]);

    const jspb = require('google-protobuf');
    const client = new PictureClient(`http://${window.location.hostname}:8080/server`, {}, {});
    const getPictures = () => {
        if (num <= 0) return;
        const request = new GetPicturesRequest();
        request.setNum(num);
        client.getPictures(request, {}, (err: Error, response: GetPicturesReply) => {
            if (err || response === null) {
                throw err;
            }
            setPictures(jspb.Message.bytesListAsB64(response.getPicturesList()).map((images: string, index: number) => (
                <img key={`${index}`} width="200px"
                    src={`data:image/jpg;base64,${window.atob(images)}`}
                    alt="pictures"
                />
            )));
        });
    }

    const getStreamPictures = () => {
        let n: number = 0;
        setPictures([]);
        const client = new PictureClient(`http://${window.location.hostname}:8080/server`, null, null);
        const sensorRequest = new GetPicturesRequest()
        sensorRequest.setNum(num);
        const stream: ClientReadableStream<StreamGetPicturesReply> = client.streamGetPictures(sensorRequest, {}) as ClientReadableStream<StreamGetPicturesReply>;
        stream.on('data', (response: StreamGetPicturesReply) => {
            setPictures(prevPictures2 => [...prevPictures2, <img key={n += 1} width="200px"
                src={`data:image/jpg;base64,${window.atob(response.getPicture_asB64())}`} alt="picture" />]
            );
        });

        stream.on('end', () => {
            console.log("END!!"); // 全て受け取った
        });
    }

    const onChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const n = event.target.valueAsNumber;
        if (!isNaN(n)) { setNumber(n); }
    };

    return (
        <div>
            <input type="number" min="1" defaultValue="1" onChange={onChange} />
            <button onClick={getPictures}>GetPictures</button>
            <button onClick={getStreamPictures}>GetStreamPictures</button>
            <div className="getPictures">
                {pictures}
            </div>
        </div>
    );
}
