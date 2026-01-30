// import Lanyard from "../team/Lanyard";

export default function MemberCard({ member }) {
  return (
    <div
      style={{
        width: "520px",
        maxWidth: "90vw",
        background: "rgba(10,25,50,0.6)",
        border: "1px solid rgba(80,160,255,0.4)",
        borderRadius: "10px",
        padding: "18px",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        gap: "16px"
      }}
    >
      {/* LANYARD */}
      {/* <div style={{ width: "100%", height: "380px" }}>
        <Lanyard />
      </div> */}

      {/* MEMBER INFO */}
      <div
        style={{
          width: "100%",
          padding: "12px",
          borderTop: "1px solid rgba(80,160,255,0.3)",
          textAlign: "center",
          color: "#cfe9ff"
        }}
      >
        <h3
          style={{
            letterSpacing: "0.2em",
            fontSize: "14px",
            marginBottom: "6px"
          }}
        >
          {member.toUpperCase()}
        </h3>

        <p
          style={{
            fontSize: "11px",
            opacity: 0.7,
            letterSpacing: "0.15em"
          }}
        >
          ACTIVE MEMBER
        </p>
      </div>
    </div>
  );
}
