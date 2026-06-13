interface RiskBadgeProps {
    score: number;
}

export default function RiskBadge({score}: RiskBadgeProps) {
    let label = "Low";

    if (score >= 80) {
        label = "Critical"
    } else if (score >= 60) {
        label = "High";
    } else if (score >= 40) {
        label = "Medium";
    }

    return (
        <span>
            {label} ({score})
        </span>
    )
}