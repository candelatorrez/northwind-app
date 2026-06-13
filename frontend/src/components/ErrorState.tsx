type Props = {
    message: string;
}

export default function ErrorState({message}: Props) {
    return (
        <h2>
            Error: {" "} {message}
        </h2>
    )
}