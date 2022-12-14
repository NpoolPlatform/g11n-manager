//nolint:nolintlint,dupl
package country

import (
	"context"

	converter "github.com/NpoolPlatform/g11n-manager/pkg/converter/country"
	crud "github.com/NpoolPlatform/g11n-manager/pkg/crud/country"
	commontracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer/country"

	constant "github.com/NpoolPlatform/g11n-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/country"

	"github.com/google/uuid"
)

func (s *Server) CreateCountry(ctx context.Context, in *npool.CreateCountryRequest) (*npool.CreateCountryResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCountry")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = validate(in.GetInfo())
	if err != nil {
		return &npool.CreateCountryResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "country", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create country: %v", err.Error())
		return &npool.CreateCountryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCountryResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateCountries(ctx context.Context, in *npool.CreateCountriesRequest) (*npool.CreateCountriesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCountries")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateCountriesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	err = duplicate(in.GetInfos())
	if err != nil {
		return &npool.CreateCountriesResponse{}, err
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "country", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create countrys: %v", err)
		return &npool.CreateCountriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCountriesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateCountry(ctx context.Context, in *npool.UpdateCountryRequest) (*npool.UpdateCountryResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateCountry")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		return &npool.UpdateCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().Country != nil && in.GetInfo().GetCountry() == "" {
		return &npool.UpdateCountryResponse{}, status.Error(codes.InvalidArgument, "Country is invalid")
	}
	if in.GetInfo().Flag != nil && in.GetInfo().GetFlag() == "" {
		return &npool.UpdateCountryResponse{}, status.Error(codes.InvalidArgument, "Flag is invalid")
	}
	if in.GetInfo().Code != nil && in.GetInfo().GetCode() == "" {
		return &npool.UpdateCountryResponse{}, status.Error(codes.InvalidArgument, "Code is invalid")
	}
	if in.GetInfo().Short != nil && in.GetInfo().GetShort() == "" {
		return &npool.UpdateCountryResponse{}, status.Error(codes.InvalidArgument, "Short is invalid")
	}

	span = commontracer.TraceInvoker(span, "country", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create country: %v", err.Error())
		return &npool.UpdateCountryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCountryResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCountry(ctx context.Context, in *npool.GetCountryRequest) (*npool.GetCountryResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCountry")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "country", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get country: %v", err)
		return &npool.GetCountryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCountryResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCountryOnly(ctx context.Context, in *npool.GetCountryOnlyRequest) (*npool.GetCountryOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCountryOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "country", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get countrys: %v", err)
		return &npool.GetCountryOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCountryOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCountries(ctx context.Context, in *npool.GetCountriesRequest) (*npool.GetCountriesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCountries")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "country", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get countrys: %v", err)
		return &npool.GetCountriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCountriesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistCountry(ctx context.Context, in *npool.ExistCountryRequest) (*npool.ExistCountryResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCountry")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "country", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check country: %v", err)
		return &npool.ExistCountryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCountryResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistCountryConds(ctx context.Context,
	in *npool.ExistCountryCondsRequest) (*npool.ExistCountryCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCountryConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "country", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check country: %v", err)
		return &npool.ExistCountryCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCountryCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountCountries(ctx context.Context, in *npool.CountCountriesRequest) (*npool.CountCountriesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountCountries")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "country", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count countrys: %v", err)
		return &npool.CountCountriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountCountriesResponse{
		Info: total,
	}, nil
}
