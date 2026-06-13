type Props = {
    title: string;
    value: string | number;
}

export default function KPICard({title, value}: Props) {
    return (
        <div>
            <div>{title}</div>
            <h2>{value}</h2>
        </div>
    )
}