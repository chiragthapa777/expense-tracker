import { Image } from "expo-image";
import React from "react";
import { StyleSheet } from "react-native";
const blurhash =
  "|rF?hV%2WCj[ayj[a|j[az_NaeWBj@ayfRayfQfQM{M|azj[azf6fQfQfQIpWXofj[ayj[j[fQayWCoeoeaya}j[ayfQa{oLj?j[WVj[ayayj[fQoff7azayj[ayj[j[ayofayayayj[fQj[ayayj[ayfjj[j[ayjuayj[";

type Props = {
  size?: number;
  imageUrl: string;
};

export default function Avatar({ imageUrl, size = 100 }: Props) {
  return (
    <Image
      style={[
        {
          height: size,
          width: size,
          borderRadius: size / 2,
          objectFit:"cover"
        },
      ]}
      source={imageUrl}
      placeholder={{ blurhash }}
      contentFit="cover"
      transition={1000}
    />
  );
}

