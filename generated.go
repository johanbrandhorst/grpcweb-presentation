type BookServiceClient interface {
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpcweb.CallOption) (*Book, error)
	QueryBooks(ctx context.Context, in *QueryBooksRequest, opts ...grpcweb.CallOption) (BookService_QueryBooksClient, error)
	MakeCollection(ctx context.Context, opts ...grpcweb.CallOption) (BookService_MakeCollectionClient, error)
	BookChat(ctx context.Context, opts ...grpcweb.CallOption) (BookService_BookChatClient, error)
}

type bookServiceClient struct {
	client *grpcweb.Client
}
func NewBookServiceClient(hostname string, opts ...grpcweb.DialOption) BookServiceClient {
	return &bookServiceClient{
		client: grpcweb.NewClient(hostname, "library.BookService", opts...),
	}
}
